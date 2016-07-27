# installation guide

## dependencies

* Cassandra. We run and recommend 3.0.8 .  We used to run 2.2.3 which was fine too. See cassandra.md
* Elasticsearch is currently a dependency for metrics metadata, but we will remove this soon.
* optionally a queue: Kafka 0.10 is reccomended, but 0.9 should work too.

## installation

### from source

Building metrictank requires a [Golang](https://golang.org/) compiler.
We recommend version 1.5 or higher.

```
go get github.com/raintank/metrictank
```

### distribution packages

#### bleeding edge packages

https://packagecloud.io/app/raintank/raintank/search?filter=all&q=metrictank&dist=

#### stable packages

TODO: stable packages, rpms

### using chef
https://github.com/raintank/chef_metric_tank

### docker

TODO

## set up cassandra

You can have metrictank initialize Cassandra with a schema without replication, good for development setups.
Or you may want to tweak the schema yourself. See schema.md

## configuration

See the [example config file](https://github.com/raintank/metrictank/blob/master/metrictank-sample.ini) which guides you through the various options


# Test with Docker.

## Install Docker
https://docs.docker.com/installation/#installation

## Start Cassandra instance
`docker run -d --name cassandra1 -p 9042:9042 cassandra:3.0.8`

## start Elasticsearch
`docker run -d --name elastic1 -p 9200:9200 elasticsearch:latest`

## start metrictank
###create metrictank.ini with 
```
listen = :6060

accounting-period = 5min

instance = default
primary-node = true
chunkspan = 30min
numchunks = 7
ttl = 35d
agg-settings = 10min:6h:2:38d:true,2h:6h:2:120d:true
cassandra-addrs = cassandra
elastic-addr = elasticsearch:9200
index-name = metric
log-level = 2


[carbon-in]
enabled = true
schemas-file = /etc/raintank/schemas.conf

```
###create schemas.conf with
```
TODO
```

### Launch container
`docker run -d --name metrictank -p 6060:6060 -p 2003:2003 --link cassandra1:cassandra --link elastic1:elasticsearch -v metrictank.ini:/etc/raintank/metrictank.ini -v schema.conf:/etc/raintank/schema.conf raintank/metrictank`


## Send metrics
send metrics from you tools like you would send to graphite using the plaintext protocol
http://graphite.readthedocs.io/en/latest/feeding-carbon.html

