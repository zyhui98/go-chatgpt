FROM centos
ENV MY_SERVICE_PORT=8080
LABEL go="1.14" version="1.0"
ADD bin/amd64/go-chatgpt /go-chatgpt
ADD configs/ /configs/
ADD html/ /html/

EXPOSE 8080
ENTRYPOINT /go-chatgpt
