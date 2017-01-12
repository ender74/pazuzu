ENV GOPATH=/usr/GOPATH
RUN mkdir -p /usr/GOPATH
RUN apt-get update
RUN apt-get install -y git
RUN rm -rf /var/lib/apt/lists/*
RUN git clone https://github.com/zalando-incubator/pazuzu.git $GOPATH/src/github.com/zalando-incubator/pazuzu
WORKDIR $GOPATH/src/github.com/zalando-incubator/pazuzu/
RUN go get -t -v
RUN go test -v ./...; exit 0
WORKDIR $GOPATH/src/github.com/zalando-incubator/pazuzu/cli/pazuzu
RUN go get -v
RUN go build
RUN go install
ENV PATH=$GOPATH/bin:$PATH