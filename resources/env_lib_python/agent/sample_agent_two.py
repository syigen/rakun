import asyncio
import logging

log = logging.getLogger("AGENT_TWO")


class AgentTwo:
    name = "Agent Sample"
    publish = None  # No need to define it will automatically assign at runtime self.publish(AGENT_NAME_OR_CLASS,<MESSAGE_OBJECT>)

    def __init__(self, *args, **kwargs):
        log.info(f"{self.name} Start")
        log.info(f"Args = {args}")
        log.info(f"Kwargs = {kwargs}")

    async def start(self):
        await self.publish("AgentOne", "Hi Agent 1")
        await self.publish("AgentOne", "Im Agent 2")

    async def accept_message(self, agent, message):
        log.info("Inbox")
        log.info(agent)
        log.info(message)

    async def stop(self, *args, **kwargs):
        log.info(f"{self.name} Stop")
        log.info(f"Args = {args}")
        log.info(f"Kwargs = {kwargs}")

    async def execute(self, *args, **kwargs):
        log.info(f"{self.name} Execute")
        log.info(f"Args = {args}")
        log.info(f"Kwargs = {kwargs}")

        while True:
            await self.publish("AgentOne", "Hi Agent 2")
            await self.publish("AgentOne", "Do you know me")
            await asyncio.sleep(2)
