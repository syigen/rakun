from agent.sample_agent_two import AgentTwo


class AgentOne:
    name = "Agent Sample"

    def __init__(self, *args, **kwargs):
        print(f"{self.name} Start")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")

    def stop(self, *args, **kwargs):
        print(f"{self.name} Stop")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")

    async def execute(self, *args, **kwargs):
        print(f"{self.name} Execute")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")
        await self.publish(AgentTwo, "Hi Agent 2")
        await self.publish(AgentTwo, "Im Agent 1")
