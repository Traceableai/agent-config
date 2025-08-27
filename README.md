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

### Updating Submodules

While updating submodules, do not add your commits inside submodule directories.
After pulling new changes in the submodule,
if you see some untracked changes in submodule, delete those untracked changes in the submodule.  
If you see some new commits in nested submodule, update the nested submodule in your local to the commit it points to in the new version of the submodule.
After these, there should not be any changes to be committed in the submodule directory.
In the root directory, add and commit the change to update submodules.

## C# Generator

```shell
protoc --csharp_out=. --plugin=protoc-gen-envloader=tools/csharp-generator/bin/Release/net10.0/osx-arm64/Traceable.Csharp.Generator  --envloader_out=. proto/ai/traceable/agent/config/v1/config.proto
```