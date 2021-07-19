############################
# STEP 1 build executable binary
############################
FROM golang:1.16-alpine  as builder
# install package for build
RUN apk -U --no-cache add \
    build-base git gcc bash tzdata git make ca-certificates \
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
RUN BUILD=${BUILDDOCKER} \
    make build-in-docker move-in-docker

############################
# STEP 2 build a small image
############################
FROM scratch

WORKDIR /app

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy our static executable
COPY --from=builder /app /app

# Set TimeZone
ENV TZ=Asia/Bangkok

ENV PATH=/app:$PATH

ENTRYPOINT ["AppMain"]