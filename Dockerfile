ARG GOVERSION=1.18
ARG ALPINEVERSION=3.15
ARG APPRUN=AppMain
############################
# STEP 1 build executable binary
############################
FROM golang:${GOVERSION}-alpine${ALPINEVERSION}  as builder
# install package for build
RUN apk -U --no-cache add \
    build-base git gcc tzdata make ca-certificates \
    && update-ca-certificates

# Set TimeZone (require)
ENV TZ=Asia/Bangkok

WORKDIR /go/src/AppBuild

# Download Modules
ADD *.mod .
RUN go mod download

# add project 
# ควรมีไฟล์ .dockerignore เพื่อไม่เอาไฟล์ที่ไม่จำเป็น
ADD . .

# Create directory for binary
RUN mkdir -p /app 

ARG BUILDDOCKER
# build go file
# Set build && move in `Makefile`
RUN make build-in-docker move-in-docker

############################
# STEP 2 build a small image
############################
FROM alpine:${ALPINEVERSION}

ARG APPRUN

WORKDIR /app

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /app /app

# Set TimeZone
ENV TZ=Asia/Bangkok
ENV APPRUN=${APPRUN}
ENV ENVOLOPMENT=production

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["sh", "-c", "/app/${APPRUN}"]