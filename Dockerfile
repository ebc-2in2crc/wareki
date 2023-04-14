FROM alpine:3.17.3

LABEL com.github.ebc-2in2crc.wareki.image=https://github.com/ebc-2in2crc/wareki.git

RUN apk update && \
	apk --no-cache add curl && \
	curl --location --remote-name https://github.com/ebc-2in2crc/wareki/releases/download/v1.2.1/wareki_linux_amd64.zip && \
	apk del curl && \
	unzip wareki_linux_amd64.zip wareki_linux_amd64/wareki && \
	cp ./wareki_linux_amd64/wareki /usr/local/bin && \
	rm -rf wareki_linux_amd64.zip wareki_linux_amd64/wareki

COPY docker-entrypoint.sh /usr/local/bin

ENTRYPOINT ["docker-entrypoint.sh"]
