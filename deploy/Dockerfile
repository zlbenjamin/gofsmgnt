# Command:
# docker build -t gofsmgnt:v0.1 .
# Note, there is a "." at the end.
FROM centos

RUN mkdir /home/app

COPY ./gofsmgnt /home/app/

EXPOSE 40000

ENTRYPOINT ["/home/app/gofsmgnt"]
