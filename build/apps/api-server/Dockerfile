FROM debian:buster-slim

LABEL maintainer="Rex Wu <tsunejui@gmail.com>"

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

ENV BIN_DIR=/api-server/bin
ENV PATH="${BIN_DIR}:${PATH}"
RUN mkdir -p ${BIN_DIR}

COPY bin/. ${BIN_DIR}/
# 'chmod +x' all executables except checksum files
RUN ls ${BIN_DIR}/* | egrep -v ".*\.(sha.+|md5)$" | xargs chmod +x

WORKDIR /api-server
