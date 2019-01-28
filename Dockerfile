# Build box
FROM golang:1.11.1 AS build

RUN mkdir -p /home/main
WORKDIR /home/main
ADD . /home/main
RUN go get -d ./...

# Build
ARG build
ARG version
ARG serviceName
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${version} -X main.Build=${build}" -o ${serviceName}
RUN cp ${serviceName} /

# Final
FROM alpine
ARG serviceName
RUN apk update
RUN apk upgrade
RUN apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN apk add curl
RUN rm -rf /var/cache/apk/*

# Move 
COPY --from=build /${serviceName} /home/

# Set TimeZone
ENV TZ=Europe/London

# Entrypoint
WORKDIR /home
ENV _SERVICENAME=${serviceName}
RUN echo "#!/bin/bash" > ./entrypoint.sh
RUN echo "./${serviceName}" >> ./entrypoint.sh
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["sh", "./entrypoint.sh"]

# Expose Port
EXPOSE 80

