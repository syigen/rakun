import asyncio
import os
import pickle

import aioredis
import click
from importlib.machinery import SourceFileLoader

from aioredis import Redis


class AgentWrapper:

    def __init__(self, id, className, source, init_params, publish):
        self.publish = publish
        agent_source_code = SourceFileLoader("", f"{os.getcwd()}/{source}").load_module()
        source_class = getattr(agent_source_code, className)
        self._agent_ = source_class(init_params=init_params)

    def stop_agent(self):
        self._agent_.stop()

    async def execute_agent(self):
        self._agent_.execute()


class PubSub:

    def __init__(self, pub: Redis = None, sub: Redis = None, channel: aioredis.Channel = None) -> None:
        self.channel = channel
        self.pub = pub
        self.sub = sub

    async def publish(self, message):
        await self.pub.publish(self.channel, pickle.dumps(message))

    async def subscribe(self, receiver):
        while await self.channel.wait_message():
            msg = await self.channel.get()
            receiver(pickle.loads(data=msg))


@click.command()
@click.option('--stack-name', help='Agent Stack Name')
@click.option('--comm-url', help='Rakun Stack Communication URL')
@click.option('--id', help='Agent ID')
@click.option('--name', help='Agent Name')
@click.option('--source', help='Agent Source')
@click.option("--init-params", multiple=True, default=[("name", "agent_init")], type=click.Tuple([str, str]))
def run(stack_name, comm_url, id, name, source, init_params):
    print(f"Agent Stack Name {stack_name}")
    print(f"Communication URL {comm_url}")
    print(f"Agent ID {id}")
    print(f"Agent Name {name}")
    print(f"Agent Source {source}")

    async def start_app():
        pub = await aioredis.create_redis(
            f'redis://{comm_url}')
        sub = await aioredis.create_redis(
            f'redis://{comm_url}')
        res = await sub.subscribe(f'{id}:1')
        ch1: aioredis.Channel = res[0]

        pub_sub = PubSub(pub=pub, sub=sub, channel=ch1)

        agent = AgentWrapper(id=id, source=source, className=name, init_params=init_params,
                             publish=pub_sub.publish)

        tsk = asyncio.ensure_future(pub_sub.subscribe(agent.execute_agent))

        # res = await pub.publish_json(f'{id}:1', ["Hello", "world"])
        # assert res == 1
        #
        # await sub.unsubscribe(f'{id}:1')
        await tsk
        sub.close()
        pub.close()

    asyncio.run(start_app())

    #
    # asyncio.run(agent.execute_agent())
    # agent.stop_agent()


if __name__ == '__main__':
    run()
