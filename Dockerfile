FROM golang:alpine
COPY config config
ADD app /
CMD ["/app"]

EXPOSE 8080