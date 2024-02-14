FROM alpine:3.8

# This Dockerfile is optimized for go binaries, change it as much as necessary
# for your language of choice.

EXPOSE 9091

COPY bin/car-pooling-challenge /

ENTRYPOINT [ "/car-pooling-challenge" ]
