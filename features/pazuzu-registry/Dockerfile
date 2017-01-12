RUN apt-get update
RUN apt-get install -y maven git
ENV PAZUZU_HOME=/usr/local/pazuzu
RUN mkdir -p $PAZUZU_HOME
RUN git clone https://github.com/zalando-incubator/pazuzu-registry.git $PAZUZU_HOME/pazuzu-registry
WORKDIR $PAZUZU_HOME/pazuzu-registry
RUN mvn clean install
EXPOSE 8080
ENTRYPOINT ["mvn", "spring-boot:run"]