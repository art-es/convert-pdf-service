FROM node:12-alpine

RUN apk add --update --no-cache \
    nodejs \
    nodejs-npm \
    openjdk8-jre-base \
    # Chromium dependencies
    nss \
    chromium-chromedriver \
    chromium

RUN sed -i -e 's/v3.11/edge/g' /etc/apk/repositories
ENV CHROME_BIN /usr/bin/chromium-browser
ENV PUPPETEER_SKIP_CHROMIUM_DOWNLOAD true
COPY ./static/fonts/* /usr/share/fonts/truetype/

CMD [ "node", "/converter/src/server.js" ]
