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
```
➜  url-shortener git:(master) ✗ k6 run load_testing.js
   
             /\      |‾‾|  /‾‾/  /‾/   
        /\  /  \     |  |_/  /  / /    
       /  \/    \    |      |  /  ‾‾\  
      /          \   |  |‾\  \ | (_) | 
     / __________ \  |__|  \__\ \___/ .io
   
     execution: local
        script: load_testing.js
        output: -
   
     scenarios: (100.00%) 1 executors, 100 max VUs, 40s max duration (incl. graceful stop):
              * default: 100 looping VUs for 10s (gracefulStop: 30s)
   
   
   running (13.8s), 000/100 VUs, 307 complete and 0 interrupted iterations
   default ✓ [======================================] 100 VUs  10s
   
   
       data_received..............: 44 MB  3.2 MB/s
       data_sent..................: 193 kB 14 kB/s
       http_req_blocked...........: avg=20.21ms  min=1.09µs   med=12.28µs  max=1.05s    p(90)=55.96ms  p(95)=125.84ms
       http_req_connecting........: avg=10.18ms  min=0s       med=0s       max=111.61ms p(90)=44.41ms  p(95)=54.89ms 
       http_req_duration..........: avg=720.94ms min=849.15µs med=44.7ms   max=7.32s    p(90)=2.98s    p(95)=4s      
       http_req_receiving.........: avg=659.74ms min=14.06µs  med=165.31µs max=7.28s    p(90)=2.85s    p(95)=3.86s   
       http_req_sending...........: avg=115.29µs min=5.57µs   med=56µs     max=2.84ms   p(90)=185.16µs p(95)=374.61µs
       http_req_tls_handshaking...: avg=9.97ms   min=0s       med=0s       max=1.01s    p(90)=0s       p(95)=83.17ms 
       http_req_waiting...........: avg=61.08ms  min=780.77µs med=43.95ms  max=771.02ms p(90)=124.57ms p(95)=294.54ms
       http_reqs..................: 1228   89.1837/s
       iteration_duration.........: avg=3.96s    min=1.35s    med=3.83s    max=8.75s    p(90)=6.38s    p(95)=7.16s   
       iterations.................: 307    22.295925/s
       vus........................: 18     min=18  max=100
       vus_max....................: 100    min=100 max=100
   
```