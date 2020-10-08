class AgentSampleOne:
    name = "Agent Sample"

    def start(self, *args, **kwargs):
        print(f"{self.name} Start")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")

    def stop(self, *args, **kwargs):
        print(f"{self.name} Stop")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")

    def execute(self, *args, **kwargs):
        print(f"{self.name} Execute")
        print(f"Args = {args}")
        print(f"Kwargs = {kwargs}")
