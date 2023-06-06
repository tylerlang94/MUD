# builds the docker container for the telnet server
# Uses golang 1.20

FROM golang:1.20
LABEL authors="tylerlang"

ENTRYPOINT ["top", "-b"]