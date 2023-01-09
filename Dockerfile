FROM golang:1.19.3-alpine3.17 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/rest-banking-api-go-tutorial

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11
COPY --from=build /go/bin/rest-banking-api-go-tutorial /
EXPOSE 3000

CMD [ "/rest-banking-api-go-tutorial" ]