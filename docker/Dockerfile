# Run on Ubuntu
FROM ubuntu:16.04

# Get arg
ARG DGRAPHIP

# Set the Workdir to weaviate
RUN mkdir -p /var/weaviate && cd /var/weaviate
WORKDIR /var/weaviate

# Install needed packages and scripts
RUN echo "BUILDING WITH DGRAPH IP: $DGRAPHIP"

RUN apt-get -qq update && apt-get -qq install -y jq curl zip wget && \
    wget -q -O /var/weaviate/weaviate.zip https://raw.githubusercontent.com/weaviate/weaviate/develop/dist/weaviate_nightly_linux_amd64.zip && \
    unzip -o -q -j weaviate.zip && \
    rm /var/weaviate/weaviate.zip && \
    rm -rf weaviate_nightly_linux_amd64.zip && \
    wget -q -O /var/weaviate/weaviate.conf.json https://raw.githubusercontent.com/weaviate/weaviate/develop/weaviate.conf.json && \
    chmod +x /var/weaviate/weaviate && \
    echo "UPDATE WITH" ${DGRAPHIP} && \
    JSONRESULT=$(jq ".environments[] | select(.name | contains(\"docker\")) | .database.database_config.host = \"${DGRAPHIP}\"" /var/weaviate/weaviate.conf.json) && \
    echo "{ \"environments\": [ $JSONRESULT ] }"  > /var/weaviate/weaviate.conf.json

# Expose dgraph ports
EXPOSE 80

# Run!
CMD /var/weaviate/weaviate --scheme=http --port=80 --host=localhost --config=docker --config-file=/var/weaviate/weaviate.conf.json
