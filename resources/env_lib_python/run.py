import asyncio
import datetime
import os
import pickle

import aioredis
import click
from importlib.machinery import SourceFileLoader

from aioredis import Redis
import logging

logging.basicConfig(level=os.environ.get("LOGLEVEL", "INFO"))
log = logging.getLogger("RAKUN-MAS")
log.info("Initiate")


class AgentWrapper:

    def __init__(self, id, agent, publish):
        self.id = id
        self.publish = publish
        self._agent_ = agent
        self._agent_.publish = publish

    async def start_agent(self):
        await self._agent_.start()

    async def stop_agent(self):
        await self._agent_.stop()

    async def execute_agent(self):
        await self._agent_.execute()

    async def accept_message(self, channel, message):
        await self._agent_.accept_message(agent=channel, message=message)


class PubSub:

    def __init__(self, pub: Redis = None, sub: Redis = None, channel: aioredis.Channel = None) -> None:
        self.channel = channel
        self.pub = pub
        self.sub = sub

    async def publish(self, agent, message):
        channel_name = agent.__name__ if type(agent) != str else agent
        log.info(f"Outgoing Message received:{datetime.datetime.now()}")
        log.info(f"Outgoing Message To Channel:{channel_name}")
        log.info(f"Outgoing Message Data:{message}")
        data = {
            "channel": channel_name,
            "message": message
        }
        await self.pub.publish(f"{channel_name}:1", pickle.dumps(data))

    async def subscribe(self, receiver):
        while await self.channel.wait_message():
            msg = await self.channel.get()
            data = pickle.loads(msg)
            sender_channel = data['channel']
            sender_message = data['message']
            log.info(f"Incoming Message received:{datetime.datetime.now()}")
            log.info(f"Incoming Message Channel:{sender_channel}")
            log.info(f"Incoming Message Data:{sender_message}")
            await receiver(sender_channel, sender_message)


@click.command()
@click.option('--stack-name', help='Agent Stack Name')
@click.option('--comm-url', help='Rakun Stack Communication URL')
@click.option('--id', help='Agent ID')
@click.option('--name', help='Agent Name')
@click.option('--source', help='Agent Source')
@click.option("--init-params", multiple=True, default=[("name", "agent_init")], type=click.Tuple([str, str]))
def run(stack_name, comm_url, id, name, source, init_params):
    log.info(f"Agent Stack Name {stack_name}")
    log.info(f"Communication URL {comm_url}")
    log.info(f"Agent ID {id}")
    log.info(f"Agent Name {name}")
    log.info(f"Agent Source {source}")

    agent_source_code = SourceFileLoader("", f"{os.getcwd()}/{source}").load_module()
    agent_class = getattr(agent_source_code, name)
    agent_obj = agent_class(init_params=init_params)

    channel_name = agent_class.__name__

    async def start_app():
        pub = await aioredis.create_redis(
            f'redis://{comm_url}')
        sub = await aioredis.create_redis(
            f'redis://{comm_url}')
        res = await sub.subscribe(f'{channel_name}:1')
        ch1: aioredis.Channel = res[0]

        pub_sub = PubSub(pub=pub, sub=sub, channel=ch1)

        agent = AgentWrapper(id=id, agent=agent_obj,
                             publish=pub_sub.publish)

        await agent.start_agent()

        # tsk = asyncio.ensure_future(pub_sub.subscribe(agent.accept_message))
        # tsk_execute = asyncio.ensure_future(pub_sub.subscribe(agent.accept_message))
        tasks = [pub_sub.subscribe(agent.accept_message), agent.execute_agent()]
        tsk = asyncio.wait(tasks, return_when=asyncio.ALL_COMPLETED)
        await tsk
        await agent.stop_agent()
        sub.close()
        pub.close()

    asyncio.run(start_app())


if __name__ == '__main__':
    run()
