
```
b>curl -v http://localhost:5000/
*   Trying ::1...
* TCP_NODELAY set
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 5000 (#0)
> GET / HTTP/1.1
> Host: localhost:5000
> User-Agent: curl/7.55.1
> Accept: */*
>
* HTTP 1.0, assume close after body
< HTTP/1.0 200 OK
< Content-Type: text/html; charset=utf-8
< Content-Length: 11
< Server: Werkzeug/2.0.1 Python/3.9.1
< Date: Tue, 01 Jun 2021 09:07:49 GMT
<
hello,world* Closing connection 0

```