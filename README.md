# CEMQ

CEMQ is an EMQX client-cli to check EMQX clusters. It provides commands to check a EMQX
cluster as a client. It is like a blackbox for monitoring EMQX behaviour.

## Commands overview

| Command | Description | Flags | SubCommands |
|:-------------:|------------------|-------|-----------------------|
| help | returns a full guide and description about CEMQ commands | - | - |
| config | manages the CEMQ cluster configuration | - | select, list, info |
| cluster | handles the cluster connection and health commands | - | connection, health |
| load | generates a load on the EMQX cluster | topics, pubs, subs, interval | - |
