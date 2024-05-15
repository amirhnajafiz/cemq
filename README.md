# CEMQ

A EMQX client-cli to test and benchmark your EMQX clusters.

## Components

- EMQX-cli
- Config
- Progress Bar

## Commands overview

| Command | Description | Flags | SubCommands |
|:-------------:|------------------|-------|-----------------------|
| help | returns a full guide and description about CEMQ commands | - | - |
| config | manages the CEMQ cluster configuration | - | select, list, info |
| cluster | handles the cluster connection and health commands | - | connection, health |
| bench | runs benchmarks on the selected EMQX cluster | topic, pub, sub, progress | - |

### Commands details

#### config

All EMQX config files are in ```.json``` format and stored in ```~/.config/cemq/configs``` directory. There is also a ```config.txt``` file which stores the current selection.
