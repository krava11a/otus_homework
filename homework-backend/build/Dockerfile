FROM scratch
COPY ./build/homework-backend /
COPY ./internal/config/config_for_docker.yaml /
EXPOSE 9001
CMD [ "/homework-backend", "-config", "/config_for_docker.yaml"]
# ENTRYPOINT ["/app/test"]