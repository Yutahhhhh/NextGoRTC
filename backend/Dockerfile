FROM golang:latest

WORKDIR /usr/src/app

RUN go install github.com/volatiletech/sqlboiler@latest
RUN go install github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

COPY . .

EXPOSE 8080
ENV BINDING=0.0.0.0
