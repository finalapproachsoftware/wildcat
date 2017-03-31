# wildcat
Tenant, user and Feature management for SAAS application

# Go Swagger

## Install

```
go get -u github.com/go-swagger/go-swagger/cmd/swagger
```

## Generate Server

```
swagger generate server -f swagger.yml -A wildcat 
```

# Run Server

```
go install ./cmd/wildcat-server/
wildcat-server --help

wildcat-server
```

http://127.0.0.1:49255/swagger.json