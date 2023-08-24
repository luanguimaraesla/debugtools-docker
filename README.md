# Debugtools

Tool kit for troubleshooting systems from inside Kubernetes

## Overview

This project contains three main components to help you to troubleshoot
issues related to kubernetes networking layer:

1. Set of well known linux tools tipically used for troubleshooting
   networking problems

2. A simple server that runs on port `5000` that can be used for
   testing networking from inside the cluster.

3. A docker-compose file that runs Grafana locally with a dashboard
   containing useful metrics about this simple server.

### Available tools

- curl
- wget
- git
- tcpdump
- inetutils-ping
- iputils-tracepath
- vim
- iproute2
- telnet
- dnsutils
- ruby
- ruby-dev
- gcc
- bpfcc-tools
- netcat
- tmux
- postgresql
- mysql-client
- make
- iperf
- redis-cli
- etcdctl
- golang
- fio
- fio-plot
- ipython

### Simple server endpoints

| Endpoint            | Description                          |
| ------------------- | ------------------------------------ |
| `/`                 | A simple testing web page            |
| `/health`           | Server's health                      |
| `/ping`             | Returns `pong` if everything is okay |
| `/post/<file-name>` | Creates a file on the server         |
| `/get/<file-name>`  | Reads a file from the server         |
| `/metrics`          | Metrics about the server             |

## Deploy on Kubernetes

1. Clone this repo

```bash
git clone https://github.com/luanguimaraesla/debugtools
cd debugtools
```

2. Copy `helm/debugtools/values.yaml` to `debugtools.values.yaml` and
   change the values you need.

3. Create resources on Kuberenetes using Helm.

```bash
helm upgrade \
	--install \
	--atomic \
	-f debugtools.values.yaml \
	-n debugtools-luan \
	--create-namespace \
	debugtools \
	./debugtools.values.yaml
```
