docker build -t proxy-server .
docker build -t server1 .
docker build -t server2 .

docker run -it -p 8080:8080 proxy-server
docker run -it ip 8001:8001 server1
docker run -it ip 8002:8002 server2

http://localhost:8080/server2/app
http://localhost:8080/server1/app
