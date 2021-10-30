FROM golang:1.16-alpine as builder

RUN apk update && apk upgrade && \
    apk --no-cache --update add git make gcc \
    libc-dev tzdata && \
    mkdir /app

WORKDIR /app

ENV TZ=Asia/Jakarta

# Private Repo
ARG repo_user
ARG repo_token
ARG repo_url
ARG repo_organization
ARG app_name
ENV app_name=$app_name

ENV repo_user=$repo_user
ENV repo_token=$repo_token
ENV repo_url=$repo_url
ENV repo_organization=$repo_organization
ENV GO111MODULE=on

RUN git config \
    --global --add \
    url."https://${repo_user}:${repo_token}@${repo_url}/".insteadOf \
    "https://${repo_url}/"

RUN export GOPRIVATE=${repo_url}/${repo_organization}/*

COPY . .

RUN go mod download

RUN go get bitbucket.org/liamstask/goose/cmd/goose

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o "./bin/go-chat-api" .

EXPOSE 8081

CMD ["./bin/go-chat-api"]