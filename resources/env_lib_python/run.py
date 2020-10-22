import asyncio
import datetime
import os
import pickle

import aioredis
import click
from importlib.machinery import SourceFileLoader

from aioredis import Redis
import logging
from rlog import RedisHandler

logging.basicConfig(level=os.environ.get("LOGLEVEL", "INFO"))
log = logging.getLogger("RAKUN-MAS")

DEBUG = True


class AgentWrapper:

    def __init__(self, id, agent, publish):
        self.id = id
        self.publish = publish
        self._agent_ = agent
        self._agent_.publish = publish

    async def start_agent(self):
        await self._agent_.start()
        try:
            await self._agent_.start()
        except Exception as e:
            log.error(e)

    async def stop_agent(self):
        await self._agent_.stop()
        try:
            await self._agent_.stop()
        except Exception as e:
            log.error(e)

    async def execute_agent(self):
        await self._agent_.execute()
        try:
            await self._agent_.execute()
        except Exception as e:
            log.error(e)

    async def accept_message(self, channel, message):
        await self._agent_.accept_message(agent=channel, message=message)
        try:
            await self._agent_.accept_message(agent=channel, message=message)
        except Exception as e:
            log.error(e)


class PubSub:

    def __init__(self, pub: Redis = None, sub: Redis = None, channel: aioredis.Channel = None) -> None:
        self.channel = channel
        self.pub = pub
        self.sub = sub

    async def publish(self, agent, message):
        channel_name = agent.__name__ if type(agent) != str else agent
        if DEBUG:
            log.info(f"Outgoing Message received:{datetime.datetime.now()}")
            log.info(f"Outgoing Message To Channel:{channel_name}")
            log.info(f"Outgoing Message Data:{message}")
        data = {
            "channel": self.channel.name.decode("utf8"),
            "message": message
        }
        while True:
            re_val = await self.pub.publish(f"{channel_name}", pickle.dumps(data))
            log.info(re_val)
            if re_val:
                break

    async def subscribe(self, receiver):
        while await self.channel.wait_message():
            msg = await self.channel.get()
            data = pickle.loads(msg)
            sender_channel = data['channel']
            sender_message = data['message']
            if DEBUG:
                log.info(f"Incoming Message received:{datetime.datetime.now()}")
                log.info(f"Incoming Message Sender:{sender_channel},Reciver :{self.channel.name}")
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
    PLATFORM_CH_NAME = f"{stack_name}_PLATFORM"
    PLATFORM_CTRL_CH_NAME = f"{stack_name}_PLATFORM_CTRL"

    START_COMMAND = f"{id}:START"

    log.addHandler(
        RedisHandler(channel=f'{stack_name}_PLATFORM_LOG', host=comm_url.split(":")[0],
                     port=int(comm_url.split(":")[1])))
    log.info("Initiate")

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
        pub = await aioredis.create_redis(f'redis://{comm_url}')
        sub = await aioredis.create_redis(f'redis://{comm_url}')
        sub_agent = await aioredis.create_redis(f'redis://{comm_url}')
        platform_ch_res = await sub.subscribe(PLATFORM_CH_NAME)
        platform_ch: aioredis.Channel = platform_ch_res[0]

        res = await sub_agent.subscribe(f'{channel_name}')
        ch1: aioredis.Channel = res[0]

        pub_sub = PubSub(pub=pub, sub=sub, channel=ch1)

        agent = AgentWrapper(id=id, agent=agent_obj,
                             publish=pub_sub.publish)

        async def start_agent():
            await agent.start_agent()
            tasks = [pub_sub.subscribe(agent.accept_message), agent.execute_agent()]
            tsk = asyncio.wait(tasks, return_when=asyncio.ALL_COMPLETED)
            await tsk
            await agent.stop_agent()

        await pub.publish(PLATFORM_CTRL_CH_NAME, f"INIT:{id}")
        # Platform handling commands
        while await platform_ch.wait_message():
            msg = await platform_ch.get(encoding="utf8")
            msg = str(msg)
            if str(msg) == START_COMMAND:
                agent_start_task = asyncio.wait([start_agent()], return_when=asyncio.ALL_COMPLETED)
                await agent_start_task
                await pub.publish(PLATFORM_CTRL_CH_NAME, f"FAILED:{id}")
        sub.close()
        pub.close()

    asyncio.run(start_app())


if __name__ == '__main__':
    run()
