FROM golang:1.19-bullseye

RUN apt-get update && apt-get install -y libuv1-dev libsqlite3-dev liblz4-dev libtool build-essential git-core autoconf
RUN git clone https://github.com/canonical/raft.git && \
    cd raft && \
    autoreconf -i && \
    ./configure --prefix=/usr && \
    make && \
    make install && \
    cd ..
RUN git clone https://github.com/canonical/dqlite.git && \
    cd dqlite && \
    autoreconf -i && \
    ./configure --prefix=/usr && \
    make && \
    make install && \
    cd ..

RUN git clone https://github.com/pharosrocks/pharosbbs.git /bbs && cd /bbs && CGO_LDFLAGS_ALLOW="-Wl,-z,now" go build -tags netgo -tags libsqlite3 -o bbsd cmd/bbsd/bbsd.go
WORKDIR /bbs
