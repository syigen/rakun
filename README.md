
# RAKUN

This ia a multi agent platform which allow to work at realtime



## Features

- Distributed Agents
- Async message passing and handling
- Fault Tolerance
- Cross platform
- Live monitoring
- Websocket enabled


  
## Installation 


### PreRequired Libraries

- https://github.com/timtadh/gopkgr

Install rakun with go

```bash 
git clone https://github.com/syigen/rakun
cd rakun
pkger -include \
      /resources/env_lib_python/run.py \
      -include \
      /resources/env_lib_python/requirements.txt \
      -include \
      /resources/env_lib_python/agent/sample_agent_one.py \
      -include \
      /resources/env_lib_python/agent/sample_agent_two.py \
      -include \
      /resources/env_lib_python/agent/__init__.py \
mkdir build
go build -o ./build
```
    
## Usage/Examples

Rakun simple example

```sh
rakun init
```

Create AgentFirst in  in to agent/first_agent.py after succefull intialze


```python
import asyncio
import logging

log = logging.getLogger("AgentFirst")

class AgentFirst:

  def __init__(*args,**kwargs):
    pass

  async def execute(self,*args,**kwargs):

    while True:
      # Log data
      log.info("Agent First Runnnig")
      # Communicate with agents
      await self.publish("AgentSecond","Iam agent first")
      # Use websocket to connect with environment 
      await self.display({"status":"Agent One Running"}) 
      await asyncio.sleep(1)


```

```yaml
name: SimpleExample
version: 1.0.0
buildversion: 1
agents:
  first:
    name: AgentFirst
    code: agent/first_agent.py
# Redis uri
communicationurl: 127.0.0.1:6379 
# Web socket uri
displayserverurl: 0.0.0.0:7979
requiredfiles:
  - requirements.txt

```

```sh
rakun run
```
## Tech Stack

**Application:** Golang

  
## Authors

- [@dewmal](https://www.github.com/dewmal)

  
