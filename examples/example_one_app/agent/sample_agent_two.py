import asyncio


class AgentTwo:
    name = "Agent Sample"

    def __init__(self, *args, **kwargs):
        print(f"{self.name} Start")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")

    async def start(self):
        from agent.sample_agent_one import AgentOne
        await self.publish(AgentOne, "Hi Agent 1")
        await self.publish(AgentOne, "Im Agent 2")

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
            await self.publish("AgentOne", "Hi Agent 1")
            await self.publish("AgentOne", "Do you know me")
            await asyncio.sleep(100)
