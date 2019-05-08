## CURL

### Using template.curl

```bash
curl -w "@template.curl" -o /dev/null -s "http://wordpress.com/"
```

Using curl with template will split response time latencies according to the following parameters

* **time_namelookup**: The time, in seconds, it took from the start until the name resolving was completed.
* **time_connect**: The time, in seconds, it took from the start until the TCP connect to the remote host (or proxy) was completed.
* **time_appconnect**: The time, in seconds, it took from the start until the SSL/SSH/etc connect/handshake to the remote host was completed
* **time_pretransfer**: The time, in seconds, it took from the start until the file transfer was just about to begin. This includes all pre-transfer commands and negotiations that are specific to the particular protocol(s) involved.
* **time_redirect**: The time, in seconds, it took for all redirection steps include name lookup, connect, pretransfer and transfer before the final transaction was started. time_redirect shows the complete execution time for multiple redirections.
* **time_starttransfer**: The time, in seconds, it took from the start until the first byte was just about to be transferred. This includes time_pretransfer and also the time the server needed to calculate the result.
* **time_total**: The total time, in seconds, that the full operation lasted. The time will be displayed with millisecond resolution

[Read more](https://netbeez.net/blog/http-transaction-timing-breakdown-with-curl/)
