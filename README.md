# CEMQ

CEMQ is an EMQX client-cli to check EMQX clusters. It provides commands to overview a EMQX
cluster as a user, like a blackbox for monitoring EMQX behaviour. This cli is useful when you want
to test your EMQX cluster as a client.

## Commands overview

In this table you can see all CEMQ commands and their usages. For more details you can install the cli and run each command with `--help` flag.

| Command | Description | Flags | SubCommands |
|:-------------:|------------------|:-----:|:---------------------:|
| help | returns a full guide and description about CEMQ commands | - | - |
| config | manages the CEMQ cluster configuration | - | select, list, info |
| cluster | handles the cluster connection and health commands | - | connection, health |
| load | generates a load on the EMQX cluster | topics, pubs, subs, interval | - |

## Install

After clonning into the project repository, run `make install`. Now you can run `cemq` cli. To make sure that your installation was successful, run `make test`.
