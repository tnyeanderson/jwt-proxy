FROM golang:1.22 as build

WORKDIR /go/src

COPY ./pkg ./pkg

WORKDIR /go/src/pkg

RUN CGO_ENABLED=0 go build -o /jwt-proxy  .

FROM nginx

WORKDIR /app

COPY --from=build /jwt-proxy /app/jwt-proxy

COPY ./nginx.conf.template /etc/nginx/templates/default.conf.template

COPY ./entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/jwt-proxy /app/entrypoint.sh

ENTRYPOINT /app/entrypoint.sh
