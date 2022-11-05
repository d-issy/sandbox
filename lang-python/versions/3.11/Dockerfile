FROM python:3.11-alpine

RUN set -x && \
    apk add --no-cache git && \
    pip install git+https://github.com/python/mypy.git && \
    pip install black

ENTRYPOINT []
CMD ["/bin/sh"]
