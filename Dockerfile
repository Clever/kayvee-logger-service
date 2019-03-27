FROM debian:stretch

COPY bin/kayvee-logger-service /usr/bin/kayvee-logger-service

WORKDIR /usr/bin

CMD ["/usr/bin/kayvee-logger-service"]
