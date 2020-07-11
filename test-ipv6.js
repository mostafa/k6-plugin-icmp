/*

This is a k6 test script that imports the k6-icmp-plugin.

*/

import { ping } from 'k6-plugin/icmp';  // import icmp plugin

export default function () {
    ping("::1")
}
