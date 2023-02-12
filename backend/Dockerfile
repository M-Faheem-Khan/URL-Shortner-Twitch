# Build Stage
FROM golang:alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o url_shortner_server

# Deploy Stage
FROM alpine
COPY . /app
COPY --from=builder /build/url_shortner_server /app/
WORKDIR /app
EXPOSE 3000
CMD [ "./url_shortner_server" ]

# EOF