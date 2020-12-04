FROM scratch

COPY cmd/cmd .

EXPOSE 8105

ENTRYPOINT [ "./cmd" ]