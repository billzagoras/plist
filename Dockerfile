
FROM golang:1.20 AS go_build

WORKDIR /plist
COPY . .

RUN make appserver

FROM debian:buster-slim

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    apt-get clean && \
    apt-get install build-essential -y

COPY .env  /env/.env

COPY --from=go_build /plist/appserver /bin/appserver

CMD ["./bin/appserver", "init"]