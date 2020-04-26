//hack/Dockerfile
FROM centos:7
ADD bin/{{ .App }} /bin/
ENTRYPOINT ["/bin/{{ .App }}"]