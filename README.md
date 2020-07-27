# Disclaimer

The [k6](https://github.com/loadimpact/k6) [plugin system](https://github.com/loadimpact/k6/issues/1353) is currently experimental. This plugin is a proof of concept, and it isn't supported by the k6 team, and may break in the future. USE IT AT YOUR OWN RISK!

---

# k6-plugin-icmp

IPv4 ping example script:

```javascript
import { check } from 'k6';
import { ping } from 'k6-plugin/icmp';  // import icmp plugin

export default function () {
    const hostname = "google.com";
    const count = 1;
    const interval = 1;
    const timeout = 1;
    const size = 8;
    const error = ping(hostname, count, interval, timeout, size);

    check(error, {
        "ping successful": err => err == undefined
    });
}
```

IPv6 ping example script:

```javascript
import { ping } from 'k6-plugin/icmp';  // import icmp plugin

export default function () {
    ping("::1"); // ping localhost with IPv6 address
}
```

Result output for IPv4 test:

```bash
$ sudo ./k6 run --vus 50 --duration 30s --plugin=icmp.so test.js

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
       vus: 2,  

  scenarios: (100.00%) 1 executors, 2 max VUs, 1m0s max duration (incl. graceful stop):
           * default: 2 looping VUs for 30s (gracefulStop: 30s)


running (0m30.1s), 0/2 VUs, 560 complete and 0 interrupted iterations
default ✓ [======================================] 2 VUs  30s


    ✓ ping successful

    checks..................: 100.00% ✓ 560 ✗ 0  
    data_received...........: 0 B     0 B/s
    data_sent...............: 0 B     0 B/s
    icmp.avg_rtt............: avg=4.53ms   min=4ms   med=4ms      max=9ms      p(90)=5ms     p(95)=5ms
    icmp.max_rtt............: avg=4.53ms   min=4ms   med=4ms      max=9ms      p(90)=5ms     p(95)=5ms
    icmp.min_rtt............: avg=4.53ms   min=4ms   med=4ms      max=9ms      p(90)=5ms     p(95)=5ms
    icmp.packets_loss.......: 0       0/s
    icmp.packets_received...: 560     18.594848/s
    icmp.packets_sent.......: 560     18.594848/s
    icmp.std_dev_rtt........: avg=0s       min=0s    med=0s       max=0s       p(90)=0s      p(95)=0s
    iteration_duration......: avg=107.29ms min=9.3ms med=110.52ms max=144.55ms p(90)=115.5ms p(95)=119.26ms
    iterations..............: 560     18.594848/s
    vus.....................: 2       min=2 max=2
    vus_max.................: 2       min=2 max=2
```
