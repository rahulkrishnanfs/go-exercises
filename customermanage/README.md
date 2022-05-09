
REST API to test CRUD Operations

QuickStat

1. Build the application as docker image
```
docker build . -t cmanaged:latest
```

Create certficate for HTTPS
```
 openssl req -new -subj "/C=UK/ST=Northampton/CN=customer.zaagpro.com" -newkey rsa:2048 -nodes -keyout server.key -out server.csr

 openssl x509 -req -days 365 -in "server.csr" -signkey server.key -out server.crt
 ```