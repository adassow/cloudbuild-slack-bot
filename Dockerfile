FROM golang AS build
COPY . /source
WORKDIR /source
RUN go test /source
RUN go build -o /source .
#RUN ls /source

#FROM gcr.io/distroless/static
FROM ubuntu
COPY --from=build /source/cloudbuild-slack-bot /
COPY --from=build /source/environment.env /
ENTRYPOINT ["/cloudbuild-slack-bot"]
