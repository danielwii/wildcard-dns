#!/bin/sh

# sed -i 's|DOMAIN|'$DOMAIN'|g' 	/etc/bind/*
# sed -i 's|URL|'$URL'|g' 		/etc/bind/*

sh -c 'go run tmpl.go'
sh -c '/usr/sbin/named -g -c /etc/bind/named.conf'
