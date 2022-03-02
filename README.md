# Agent config
This repo is for agents that need to configure libtraceable.

## Generate file

### Instrumentation Agents with Libtraceable
Instrumentation agents that use libtraceable should expose the ability to configure libtraceable attributes.
This can be done by using the configuration objects defined in: `ai/traceable/agent/config/v1/config.proto`

## Development

### Checking out hypertrace agent-config submodule
Remember to check out hypertrace agent-config submodule when you first clone this repo.
```
git submodule update --init --recursive
```

### Making changes to config.proto
After making changes to config.proto, run
```
make generate
```
