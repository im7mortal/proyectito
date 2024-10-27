# Created by Petr Lozhkin

FROM registry.hub.docker.com/library/golang:1.23.2-bookworm as build

ENV PROJECT_DIRECTORY=/go/src/github.com/im7mortal/project/

# Copy src to image
WORKDIR ${PROJECT_DIRECTORY}
COPY . .
WORKDIR /go/src/github.com/im7mortal/project/cmd/api

RUN go install

FROM gcr.io/distroless/base-debian12 as runtime

COPY --from=build /go/bin/api /api
COPY public /public

CMD ["/api"]
