FROM alpine:3.4

RUN apk add --no-cache bind go

# COPY bind /etc/bind
ADD bootstrap.sh .
ADD tmpl.go .

EXPOSE 53

# CMD ["sh bootstrap.sh", "&&", "/usr/sbin/named", "-g", "-c", "/etc/bind/named.conf"]
# ENTRYPOINT ["go", "run", "tmpl.go"]

CMD ["sh", "bootstrap.sh"]
