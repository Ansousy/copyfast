From alpine:latest

COPY ./build/linux/copyfast /bin/copyfast

ENTRYPOINT [ "/bin/copyfast" ]