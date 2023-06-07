FROM golang:1.19-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o publicresolvers .
RUN chmod a+x ./publicresolvers
ENTRYPOINT ["/app/publicresolvers"]