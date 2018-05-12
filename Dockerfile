FROM alpine
ADD build/server /
EXPOSE 8080
CMD ["/server"]