FROM ubuntu:latest

# Install debug tools
RUN apt update -y && apt install -y \
    curl \
    git \
    inetutils-ping \
    iputils-tracepath \
    vim \
    iproute2 \
    telnet \
    dnsutils \
    ruby \
    ruby-dev \
    netcat && \
    rm -rf /var/lib/apt/lists/*

# Install etcdctl
RUN curl -L https://github.com/coreos/etcd/releases/download/v3.3.1/etcd-v3.3.1-linux-amd64.tar.gz -o etcd-v3.3.1-linux-amd64.tar.gz && \
    tar xzvf etcd-v3.3.1-linux-amd64.tar.gz && \
    cd etcd-v3.3.1-linux-amd64 && \
    cp etcd /usr/local/bin/ && \
    cp etcdctl /usr/local/bin/ && \
    cd ../ && rm -rf etcd-v3.3.1-linux-amd64.tar.gz
ENV ETCDCTL_API=3

# Install golang
RUN curl -LO https://dl.google.com/go/go1.12.2.linux-amd64.tar.gz && \
    tar -xvf go1.12.2.linux-amd64.tar.gz && \
    rm -rf go1.12.2.linux-amd64.tar.gz && \
    mv go /usr/local && \
    mkdir -p /go/src /go/bin /go/pkg
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH

# Install vegeta
RUN go get -u github.com/tsenart/vegeta

# Configure simple server
ADD ./server /go/src/github.com/luanguimaraesla/debugtools/server

CMD ["sleep", "infinity"]