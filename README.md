

### Run the proxy server in local machine:
Run each server individually using the following command from each server directory:
```bash
go run main.go
```


### Run the proxy server in docker machine:
Docker build and run:

```bash
cd server1
docker build -t server1 .
docker run -it -p 8001:8001 server1
```

```bash
cd server2
docker build -t server2 .
docker run -it -p 8002:8002 server2
```

```bash
cd proxyServer
docker build -t proxy-server .
docker run -it -p 8080:8080 proxy-server
```