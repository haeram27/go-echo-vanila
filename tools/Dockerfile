FROM golang:1.18-alpine AS build
#FROM golang:1.18 AS build

COPY ./echoinit /tmp/echoinit
WORKDIR /tmp/echoinit
RUN go mod tidy
RUN go build -o /tmp/echoinit/echoinit

FROM alpine:latest
#FROM ubuntu:22.04
#RUN apt update
#RUN apt upgrade -y

COPY --from=build /tmp/echoinit/echoinit /usr/bin/echoinit
ENTRYPOINT ["/usr/bin/echoinit"]
#COPY --from=build /tmp/echoinit/dosleep.sh /usr/bin/dosleep.sh
#RUN chmod +x /usr/bin/dosleep.sh
#ENTRYPOINT ["/usr/bin/dosleep.sh"]
