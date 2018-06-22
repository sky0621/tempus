# step 1. compile
FROM "sky0621dhub/dockerfile-gowithdep" AS builder

COPY . /go/src/github.com/sky0621/tempus
WORKDIR /go/src/github.com/sky0621/tempus
RUN dep ensure
RUN go test ./...
RUN CGO_ENABLED=0 go build -o tempus github.com/sky0621/tempus

# -----------------------------------------------------------------------------
# step 2. build
FROM scratch
COPY --from=builder /go/src/github.com/sky0621/tempus/ .
ENTRYPOINT [ "./tempus" ]
