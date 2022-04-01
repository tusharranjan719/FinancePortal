FROM golang:1.13-alpine as builder

WORKDIR /bill_split

COPY . .

RUN go build -o bill-split


FROM alpine:latest

WORKDIR /bill_split

COPY --from=builder /bill_split/bill-split ./bill-split

EXPOSE 8010

CMD ["bill-split"]
