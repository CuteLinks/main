FROM alpine:latest
COPY . cutelinks/
CMD cutelinks/cutelinks --port 82 --host '0.0.0.0'