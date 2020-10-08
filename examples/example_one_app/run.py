import click


async def start_agent():
    pass


async def stop_agent():
    pass


async def execute_agent():
    pass


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


if __name__ == '__main__':
    run()
