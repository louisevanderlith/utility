FROM golang:1.13 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY cmd/main.go .
COPY handles ./handles
COPY core ./core

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/utility .

EXPOSE 8088

ENTRYPOINT [ "./utility" ]
