FROM mcr.microsoft.com/devcontainers/go:1.22-bookworm

# The base image's Yarn repository has a stale signing key, which breaks the
# Docker-in-Docker feature's apt-get update during Dev Container creation.
RUN rm -f /etc/apt/sources.list.d/yarn.list /etc/apt/sources.list.d/yarn.sources
