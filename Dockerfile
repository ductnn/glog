FROM golang:latest as builder
LABEL maintainer="ductnn"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o glog .


FROM alpine:latest
LABEL maintainer="ductnn"

WORKDIR /root/
COPY --from=builder /app/glog .

CMD ["./glog"]
