# =================================================================================
#  Debug docker env
# =================================================================================
FROM golang:1.13.5-stretch

ARG ARCH
ARG GET_CMD
ARG RUNTIME_DEPENDS

ENV ARCH=amd64 \
    GO111MODULE=on
