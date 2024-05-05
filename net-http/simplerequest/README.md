Ответ от сервера:

```json
{
  "args": {},
  "headers": {
    "Accept-Encoding": "gzip",
    "Host": "httpbin.org",
    "User-Agent": "Go-http-client/2.0",
    "X-Amzn-Trace-Id": "Root=1-663466d2-2f911fac409e14aa45028104"
  },
  "origin": "163.172.215.231",
  "url": "https://httpbin.org/get"
}
```


Ответ сервера с выключенной компрессией данных:
```json
{
  "args": {}, 
  "headers": {
    "Host": "httpbin.org", 
    "User-Agent": "Go-http-client/2.0", 
    "X-Amzn-Trace-Id": "Root=1-66346a1e-78ece45b216e64d930ad0be3"
  }, 
  "origin": "163.172.215.231", 
  "url": "https://httpbin.org/get"
}
```