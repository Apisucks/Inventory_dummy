FROM ubuntu:latest
WORKDIR /app
COPY ./build/service /app/bin
COPY ./build/config.json /app/config.json
ENTRYPOINT ["/app/bin"]