FROM gliderlabs/alpine:3.1

MAINTAINER Knut Ahlers <knut@ahlers.me>

RUN apk --update add wget && \
    wget --no-check-certificate https://gobuilder.me/get/github.com/Luzifer/simpleproxy/simpleproxy_master_linux-amd64.zip && \
    unzip simpleproxy_master_linux-amd64.zip

ENTRYPOINT ["/simpleproxy/simpleproxy"]
CMD ["--help"]
