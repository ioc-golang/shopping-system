FROM alpine:3.15
ARG appname

WORKDIR /shopping-system

COPY ./cmd/$appname/app .
COPY ./configs .
COPY ./configs/data ./configs/data

CMD /shopping-system/app
