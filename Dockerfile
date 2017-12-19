FROM golang
MAINTAINER Mark Chmarny <mchmarny@google.com>

RUN mkdir /app
COPY ./custom-metrics /app/custom-metrics

WORKDIR /app
CMD /app/custom-metrics
