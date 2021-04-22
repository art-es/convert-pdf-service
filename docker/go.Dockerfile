FROM golang:1.15-alpine

RUN apk add --update --no-cache \
    nodejs \
    nodejs-npm \
    openjdk8-jre-base \
    # Chromium dependencies
    nss \
    chromium-chromedriver \
    chromium

COPY ./converter /converter
WORKDIR /converter
RUN sed -i -e 's/v3.11/edge/g' /etc/apk/repositories
ENV CHROME_BIN /usr/bin/chromium-browser
ENV PUPPETEER_SKIP_CHROMIUM_DOWNLOAD true
COPY ./static/fonts/* /usr/share/fonts/truetype/
RUN npm install

WORKDIR /app

CMD [ "go", "run", "." ]
