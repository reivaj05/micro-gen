FROM python:alpine

RUN set -x; echo http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories

ADD . /mockService
WORKDIR /mockService

RUN apk -U add gcc g++ bash git python-dev libxml2-dev libxslt-dev libffi-dev openssl-dev linux-headers
RUN pip install -r requirements.txt
RUN chmod +x /mockService/scripts/*.sh
RUN /mockService/scripts/linter.sh
RUN /mockService/scripts/tests.sh
RUN apk del libxml2-dev libffi-dev openssl-dev linux-headers

EXPOSE 8000

ENTRYPOINT ./scripts/start.sh