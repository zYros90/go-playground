FROM golang:1.18.3-alpine AS builder

WORKDIR /build
COPY . .
RUN go build 

FROM alpine
WORKDIR /src
COPY --from=builder /build/http-server .

CMD [ "./http-server", "--port", "80"]