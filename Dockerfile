FROM golang:latest AS builder
ADD . /app
WORKDIR /app
RUN make deps
RUN make bin
RUN ls -aril

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder app/main ./
RUN chmod +x ./main
EXPOSE 5000
CMD ./main