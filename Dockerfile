################################################################################
# BUILDER IMAGE
################################################################################

FROM golang:1.17-alpine as builder

RUN apk add --no-cache git 

# Disable CGO as its not supported in alpine
ENV CGO_ENABLED 0

# Copy all source code and required files into the build directory
COPY . /build/
WORKDIR /build/

# Build the executable
RUN go build -o advent-of-code

################################################################################
# FINAL IMAGE
################################################################################

FROM alpine:3.14

WORKDIR /app
COPY --from=builder /build/advent-of-code ./

ENTRYPOINT [ "./advent-of-code" ]
### Note does not currently work