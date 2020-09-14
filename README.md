# url-shortener

Service for shortening URLs (like bit.ly).

### How to start server

`make`

### How to stop server

`make stop`

### How to run tests

`make tests`

### Interaction examples

1. Shorten a link:

    `curl localhost:8100/shorten -X POST -d '{"url": "https://github.com"}'`
    
    Possible response:
    `{"url":"rfBd56"}`

2. Open shortened link and redirect to `github.com`:

    `curl localhost:8100/rfBd56`

### Load testing

`k6 run load_testing.js`

Results:
```running (26.8s), 0000/1000 VUs, 1012 complete and 0 interrupted iterations
   default ✓ [======================================] 1000 VUs  10s
   
   
       data_received..............: 145 kB 5.4 kB/s
       data_sent..................: 273 kB 10 kB/s
       http_req_blocked...........: avg=1.88s    min=1.07µs  med=1.18s    max=6.66s  p(90)=4.52s    p(95)=4.68s   
       http_req_connecting........: avg=132.35ms min=0s      med=0s       max=2.29s  p(90)=211.98ms p(95)=325.53ms
       http_req_duration..........: avg=8.52s    min=70.7ms  med=4.42s    max=18.61s p(90)=18.2s    p(95)=18.41s  
       http_req_receiving.........: avg=11.61ms  min=0s      med=16.35µs  max=1.68s  p(90)=63.73µs  p(95)=345.81µs
       http_req_sending...........: avg=220.86ms min=5.12µs  med=224.72µs max=18.61s p(90)=58.86ms  p(95)=127.09ms
       http_req_tls_handshaking...: avg=0s       min=0s      med=0s       max=0s     p(90)=0s       p(95)=0s      
       http_req_waiting...........: avg=8.29s    min=48.06µs med=2.57s    max=18.61s p(90)=18.2s    p(95)=18.39s  
       http_reqs..................: 2024   75.558575/s
       iteration_duration.........: avg=20.94s   min=7.09s   med=20.9s    max=25.99s p(90)=23.53s   p(95)=23.7s   
       iterations.................: 1012   37.779288/s
       vus........................: 3      min=3    max=1000
       vus_max....................: 1000   min=1000 max=1000```