
# RAKUN

A Multi Agent system (MAS) is an intelligent agent based system which allows to solve complex problems which cannot be solved by monolithic systems. Developing these systems correctly is not an easy task. Rakun is a distributed platform which enables developers to develop MAS without worrying about the complexity of the developing MAS each time. It can be used with any existing solution without worrying about dependencies. 


## Features

- [x] Distributed Agents 
- [x] Async message passing and handling
- [x] Fault Tolerance
- [x] Cross platform
- [x] Live monitoring
- [x] Websocket enabled
- [ ] Drag and drop agent development
- [ ] Web dashboard

  
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
rakun config file
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
- [@sadika](https://www.github.com/sadika9)

  
