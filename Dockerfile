FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

# Install debug tools
RUN apt update -y && apt install -y \
    curl \
    wget \
    git \
    tcpdump \
    inetutils-ping \
    iputils-tracepath \
    vim \
    iproute2 \
    telnet \
    dnsutils \
    ruby \
    ruby-dev \
    gcc \
    nano \
    bpfcc-tools \
    netcat \
    tmux \
    postgresql \
    postgresql-contrib \
    mysql-client \
    libjemalloc-dev \
    make \
    iperf \
    fio \
    python3-pip && \
    rm -rf /var/lib/apt/lists/*

# Install redis-cli
RUN cd /tmp && \
    wget http://download.redis.io/redis-stable.tar.gz && \
    tar xvzf redis-stable.tar.gz && \
    cd redis-stable && \
    make && \
    cp src/redis-cli /usr/local/bin/ && \
    cd && rm -rf /tmp/redis-stable* && \
    chmod 755 /usr/local/bin/redis-cli

# Install etcdctl
RUN curl -L https://github.com/coreos/etcd/releases/download/v3.3.1/etcd-v3.3.1-linux-amd64.tar.gz -o etcd-v3.3.1-linux-amd64.tar.gz && \
    tar xzvf etcd-v3.3.1-linux-amd64.tar.gz && \
    cd etcd-v3.3.1-linux-amd64 && \
    cp etcd /usr/local/bin/ && \
    cp etcdctl /usr/local/bin/ && \
    cd ../ && rm -rf etcd-v3.3.1-linux-amd64.tar.gz
ENV ETCDCTL_API=3

# Install golang
RUN curl -LO https://go.dev/dl/go1.24.2.linux-amd64.tar.gz && \
    tar -xvf go1.24.2.linux-amd64.tar.gz && \
    rm -rf go1.24.2.linux-amd64.tar.gz && \
    mv go /usr/local && \
    mkdir -p /go/src /go/bin /go/pkg
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Install pip packages
RUN pip3 install \
	ipython \
	fio-plot

# Configure simple server
ADD ./server /server
WORKDIR /server
RUN go build && go install

# Release warzone tools
ADD ./warzone /warzone
WORKDIR /warzone

CMD ["server"]
