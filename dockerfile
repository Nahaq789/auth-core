FROM golang:1.23-alpine AS build_base
RUN apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -o bootstrap ./cmd/api
FROM alpine:3.9 
RUN apk add ca-certificates
COPY --from=public.ecr.aws/awsguru/aws-lambda-adapter:0.9.0 /lambda-adapter /opt/extensions/lambda-adapter
COPY --from=build_base /app/bootstrap /app/bootstrap

ENV PORT=8080 GIN_MODE=release
EXPOSE 8080

CMD ["/app/bootstrap"]
