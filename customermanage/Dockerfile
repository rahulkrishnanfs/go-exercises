From golang:1.18.1-alpine3.14 as builder
WORKDIR /customermanage
COPY . .
RUN go mod download && go build -o app ./cmanaged

FROM alpine:3.11.3
WORKDIR /cmanaged
COPY --from=builder /customermanage/app .
EXPOSE 8080
ENTRYPOINT [ "./app" ]
