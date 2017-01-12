RUN apt-get update
RUN apt-get install -y wget
RUN wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz -O /tmp/go.tar.gz
RUN tar -zxvf /tmp/go.tar.gz -C /usr/local
RUN rm /tmp/go.tar.gz
ENV PATH=/usr/local/go/bin:$PATH