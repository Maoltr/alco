FROM debian:latest

LABEL maintainer Maksym Tretiak <mmtretiak@gmail.com>

ENV PATH_TO_CONFIG="/api/config/config.json"

# copy config files
RUN mkdir -p /api/config/
COPY ./config /api/config

# copy executable file
COPY ./api /usr/bin/
RUN chmod +x /usr/bin/api

# run api
ENTRYPOINT ["api"]
