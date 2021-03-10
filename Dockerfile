FROM alpine
#FROM scratch
ADD ./ /
ENTRYPOINT ["/ze-redis.so"]