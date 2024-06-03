# CEMQ

CEMQ is an EMQX client-cli to check EMQX clusters as a client. It provides commands to overview an EMQX
cluster as a user, like a blackbox for monitoring EMQX clusters behaviour. This cli is useful when you want to test your EMQX cluster as a normal user.

CEMQ is implemented by `Golang` and `Cobra-cli`. Its simple logic avoids any bias in using an EMQX cluster
as a user.

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

## Docker

You can use `ghcr.io/amirhnajafiz/cemq:latest` image in order to run CEMQ inside a container or a pod.
An example can be found below:

```sh
docker run \
    --entrypoint cemq load --pubs 1 --subs 1 --topics 1 --server tcp://0.0.0.0:1883 \
    ghcr.io/amirhnajafiz/cemq:latest
```
