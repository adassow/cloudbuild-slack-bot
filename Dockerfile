FROM golang AS build
COPY . /source
WORKDIR /source
RUN go test /source
RUN go build -o cb-slack-bot

FROM gcr.io/distroless/base
COPY --from=build /source/cb-slack-bot /
COPY --from=build /source/environment.env /
ENTRYPOINT ["/cb-slack-bot"]
