FROM golang:1.19-alpine AS build

WORKDIR /app

COPY ./go.mod ./
RUN go mod download

COPY . .

RUN go build -o /test-app .

# Deploy
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=build /test-app /test-app

EXPOSE 8000
CMD [ "/test-app" ]