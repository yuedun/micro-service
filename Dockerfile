FROM alpine
ADD micro-srv /micro-srv
ENTRYPOINT [ "/micro-srv" ]
