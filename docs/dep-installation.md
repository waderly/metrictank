# Dependency Installation

## Cassandra

## Elasticsearch

## Kafka

## graphite-metrictank

The easiest way to install graphite-metrictank (along with our graphite-api fork) is [by adding our packagecloud repo](https://packagecloud.io/raintank/raintank/install) and installing from there.

If, for whatever reason, you can't do that:

* Install the build dependencies. 
  * Under debian based distros, run `apt-get -y install python python-pip build-essential python-dev libffi-dev libcairo2-dev git` as root. 
  * For CentOS and other rpm-based distros, run `yum -y install python-setuptools python-devel gcc gcc-c++ make openssl-devel libffi-devel cairo-devel git; easy_install pip`.
  * If neither of these instructions are relevant to you, figure out how your distribution or operating system refers to the above packages and install them.

* Install `virtualenv`, if desired: `pip install virtualenv virtualenv-tools`

* 
