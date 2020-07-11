# k6-plugin-icmp

Example script:

```javascript
import { ping } from 'k6-plugin/icmp';  // import icmp plugin

export default function () {
    ping("google.com")
}
```

Result output:

```bash
$ ./k6 run --vus 50 --duration 30s --plugin=icmp.so test.js

          /\      |‾‾|  /‾‾/  /‾/
     /\  /  \     |  |_/  /  / /
    /  \/    \    |      |  /  ‾‾\  
   /          \   |  |‾\  \ | (_) |
  / __________ \  |__|  \__\ \___/ .io

  execution: local
    plugins: ICMP
     output: -
     script: test.js

    duration: 30s, iterations: -
         vus: 50,

  execution: local
     script: test.js
     output: -

  scenarios: (100.00%) 1 executors, 50 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 50 looping VUs for 30s (gracefulStop: 30s)


running (0m30.1s), 00/50 VUs, 13381 complete and 0 interrupted iterations
default ✓ [======================================] 50 VUs  30s


    data_received...........: 0 B   0 B/s
    data_sent...............: 0 B   0 B/s
    icmp.avg_rtt............: avg=5.02µs   min=2.61µs med=4.75µs   max=39.76µs  p(90)=5.88µs   p(95)=6.5µs
    icmp.max_rtt............: avg=5.02µs   min=2.61µs med=4.75µs   max=39.76µs  p(90)=5.88µs   p(95)=6.5µs
    icmp.min_rtt............: avg=5.02µs   min=2.61µs med=4.75µs   max=39.76µs  p(90)=5.88µs   p(95)=6.5µs
    icmp.packets_loss.......: 0     0/s
    icmp.packets_received...: 13381 444.014009/s
    icmp.packets_sent.......: 13381 444.014009/s
    icmp.std_dev_rtt........: avg=0s       min=0s     med=0s       max=0s       p(90)=0s       p(95)=0s
    iteration_duration......: avg=112.13ms min=6.85ms med=110.81ms max=265.01ms p(90)=117.26ms p(95)=121.51ms
    iterations..............: 13381 444.014009/s
    vus.....................: 50    min=50 max=50
    vus_max.................: 50    min=50 max=50
```
