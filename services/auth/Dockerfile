FROM golang:1.22rc1-bookworm

WORKDIR /go/src/app

ENV MYSQL_DB=auth
ENV MYSQL_PORT=3306
ENV MYSQL_USER=test
ENV MYSQL_HOST=localhost
ENV MYSQL_PASSWORD=test
ENV JWT_SECRET=banana

COPY . .

RUN go mod download

RUN go build -o auth_server server.go


EXPOSE 3000

CMD ["./auth_server"]



