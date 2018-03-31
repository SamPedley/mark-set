FROM golang:alpine 
ADD app /
CMD ["/app"]

EXPOSE 8080