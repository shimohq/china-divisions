FROM registry-vpc.cn-beijing.aliyuncs.com/shimopro/golang:1.11 AS build-env
# FROM golang:1.11 AS build-env test

ENV GO111MODULE on
ENV HTTP_PROXY=http://home.lqs.me:23128
ENV HTTPS_PROXY=http://home.lqs.me:23128
ENV NO_PROXY=192.168.0.0/16,git.shimo.im


ADD . /data/
WORKDIR /data

RUN go mod download && go build -v .

# final image.
FROM registry-vpc.cn-beijing.aliyuncs.com/shimobase/alpine:3.8 as final

COPY --from=build-env /data/svc-china-divisions /data/svc-china-divisions
COPY --from=build-env /data/district_data/ /data/district_data/
WORKDIR /data
ENV GIN_MODE release
EXPOSE 9001

CMD /data/svc-china-divisions