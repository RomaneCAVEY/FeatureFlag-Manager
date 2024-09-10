
FROM golang:1.21.2

WORKDIR /app
ADD . /app


ENV DATABASE_HOST=db
ENV DATABASE_USER=postgres
ENV DATABASE_PASS=docker
ENV DATABASE_NAME=feature_flag

RUN go mod download

RUN go env -w GO111MODULE=on

RUN go build -o /feature_flago ./cmd/app/main.go


EXPOSE 8099

CMD [ "/feature_flago" ] 