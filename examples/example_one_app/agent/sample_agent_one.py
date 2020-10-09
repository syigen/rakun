import asyncio


class AgentOne:
    name = "Agent Sample"

    def __init__(self, *args, **kwargs):
        print(f"{self.name} Start")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")

    async def start(self):
        await self.publish("AgentTwo", "Hi Agent 2")
        await self.publish("AgentTwo", "Im Agent 1")

    async def accept_message(self, agent, message):
        print("Inbox")
        print(agent)
        print(message)

    async def stop(self, *args, **kwargs):
        print(f"{self.name} Stop")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")

    async def execute(self, *args, **kwargs):
        print(f"{self.name} Execute")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")

        while True:
            await self.publish("AgentTwo", "Hi Agent 2")
            await self.publish("AgentTwo", "Do you know me")
            await asyncio.sleep(2)
