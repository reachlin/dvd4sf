FROM alpine

RUN apk update

RUN mkdir -p /run/docker/plugins /mnt/state /mnt/volumes

COPY dvd4sf dvd4sf

CMD ["dvd4sf"]
