FROM golang:1.21-alpine as main
LABEL author="ERIK_PETROSYAN"
LABEL description="shtem-api"
RUN addgroup --system erik && adduser --system --ingroup erik erik
COPY --chown=erik:erik ./build/api/ /home/erik/cmd/api/
WORKDIR /home/erik/
RUN chmod +x ./cmd/api/app
USER erik
EXPOSE 9998