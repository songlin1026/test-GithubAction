FROM alpine:latest

RUN apk add --no-cache dumb-init curl tzdata

EXPOSE 3001

ADD ./dist /root/dist
ADD test-githubaction /root

ENV CONNECTIONSTRING=root:abcd@1234@tcp(192.168.1.160:33306)/test-githubaction?charset=utf8mb4&parseTime=True&loc=Local
ENV TZ=Asia/Taipei

WORKDIR "/root"

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ./test-githubaction

HEALTHCHECK --interval=1m --timeout=1s CMD curl -f http://localhost:3001 || exit 1