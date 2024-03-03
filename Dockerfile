FROM debian:stable-slim

COPY /build/url-shortener /bin/url-shortener

COPY config/local.yaml config/local.yaml

CMD ["/bin/url-shortener"]
