FROM centos:6.7
WORKDIR /app
COPY gobed /app/gobed
COPY app.conf /app/app.conf
COPY static/ /app/static/
COPY templates /app/templates
ENTRYPOINT ./gobed
EXPOSE 7429
