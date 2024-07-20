FROM scratch
COPY ./build/homework-dialogs /
COPY ./internal/config/config_docker.yaml /config.yaml
EXPOSE 9002
EXPOSE 9088
EXPOSE 9087
CMD [ "/homework-dialogs", "-config", "/config.yaml"]
# ENTRYPOINT ["/app/test"]