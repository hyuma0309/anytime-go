FROM golang:latest

COPY src/api /go/src/api

WORKDIR /go/src/api/

# Go
RUN go get -u github.com/gin-gonic/gin && \
    go get -u github.com/jinzhu/gorm && \
    go get -u github.com/go-sql-driver/mysql && \
    go get -u github.com/sclevine/agouti 

# Basic
RUN apt-get update \
    && apt-get install -y sudo curl wget zip unzip git

# Chrome
RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | sudo apt-key add - \
    && sh -c 'echo "deb http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list' \
    && apt-get update \
    && apt-get install -y google-chrome-beta

# ChromeDriver
RUN CHROMEDRIVER_VERSION=`curl -sS chromedriver.storage.googleapis.com/LATEST_RELEASE` \
    && mkdir -p /opt/chromedriver-$CHROMEDRIVER_VERSION \
    && curl -sS -o /tmp/chromedriver_linux64.zip http://chromedriver.storage.googleapis.com/$CHROMEDRIVER_VERSION/chromedriver_linux64.zip \
    && unzip -qq /tmp/chromedriver_linux64.zip -d /opt/chromedriver-$CHROMEDRIVER_VERSION \
    && rm /tmp/chromedriver_linux64.zip \
    && chmod +x /opt/chromedriver-$CHROMEDRIVER_VERSION/chromedriver \
    && ln -fs /opt/chromedriver-$CHROMEDRIVER_VERSION/chromedriver /usr/local/bin/chromedriver

RUN /etc/init.d/dbus start

CMD ["sh", "./start_app.sh"] 