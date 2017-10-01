FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/LucaTony/GoTest

ENV USER luca
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET lxQXJBQcnPsxgqzg

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://luca@localhost:5432/GoTest?sslmode=disable

WORKDIR /go/src/github.com/LucaTony/GoTest

RUN godep go build

EXPOSE 8888
CMD ./GoTest