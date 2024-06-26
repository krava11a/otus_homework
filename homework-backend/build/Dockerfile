FROM scratch
COPY ./build/homework-backend /
COPY ./internal/config/config_for_docker.yaml /
EXPOSE 9001
EXPOSE 8088
EXPOSE 8087
CMD [ "/homework-backend", "-config", "/config_for_docker.yaml"]
# ENTRYPOINT ["/app/test"]