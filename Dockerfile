FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY bin/app .
COPY resources ./resources

EXPOSE 9999

CMD ["./app"]