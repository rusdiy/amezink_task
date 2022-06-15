# BUILDING THE APP
FROM golang:latest AS builder
# set the current Working Directory inside the container
RUN mkdir /app
WORKDIR /app
# download dependencies if any
COPY go.mod ./
COPY go.sum ./
RUN go mod download
# copy the remaining files, and the build the app
COPY . ./
RUN go build -o /amezink_task


# DEPLOYING
FROM gcr.io/distroless/base-debian10
# copy the already-built binary from the builder, then run it
WORKDIR /
COPY --from=builder /amezink_task /amezink_task
EXPOSE 8080
ENTRYPOINT ["/amezink_task"]