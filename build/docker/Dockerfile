FROM golang:latest as build

WORKDIR /service
ADD . /service

RUN CD /service && CGO_ENABLED=0 GOOS=linux go build -o /http-service .

CMD /http-service

#test
FROM build as test

# production
FROM alpine:latest as production

RUN apk --no-cache add ca-certificates
COPY --from=build /http-service ./
RUN chmod +x ./http-service

ENTRYPOINT ["./http-service"]