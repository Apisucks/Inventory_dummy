FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /app
COPY ./build/service /app/bin
COPY ./build/config.json /app/config.json
ENTRYPOINT ["/app/bin"]
# CMD ["/app/bin"]