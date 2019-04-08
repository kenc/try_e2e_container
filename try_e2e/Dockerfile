From golang:1.12.1-alpine3.9

## install chromium
## see: https://qiita.com/dd511805/items/dfe03c5486bf1421875a

RUN apk add --update \
            udev \
            ttf-freefont \
            chromium \ 
            git

RUN mkdir /noto

ADD https://noto-website.storage.googleapis.com/pkgs/NotoSansCJKjp-hinted.zip /noto 

WORKDIR /noto

RUN unzip NotoSansCJKjp-hinted.zip && \
    mkdir -p /usr/share/fonts/noto && \
    cp *.otf /usr/share/fonts/noto && \
    chmod 644 -R /usr/share/fonts/noto/ && \
    fc-cache -fv

WORKDIR /
RUN rm -rf /noto

## install chrome driver

RUN apk add chromium-chromedriver 

 ## install dep

 RUN go get -u github.com/golang/dep/cmd/dep
