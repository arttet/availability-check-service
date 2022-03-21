# Requirements

Write a daemon that every 30 secs gets the host+port list from the MySQL database and performs service availability checks against each record from that list. The checks must be performed independently from each other. The daemon must perform 3 attempts of service availability check for each host+port and after then update the status and fail_message fields for the current record in the database. Check considered to be ok when 2 checks from the 3 were succeeded.

Service availability check - is a function that checks for the open TCP port of the host.

---

##  **checks** table with random data

```
+--------------------------------------------------+------------------+
| id | host            | port  | status | timeout  | fail_message     |
+----+-----------------+-------+--------+----------+------------------+
| 1  | google.com      | 443   | fail   | 1000     | no route to host |
| 2  | 12.12.12.12     | 80    | fail   | 5000     | timeout          |
| 3  | localhost       | 22    | ok     | 100      |                  |
+----+-----------------+-------+--------+----------+------------------+

id           - consequent number of the record
host         - host that the check performed against
port         - tcp port of the host the check performed against
status       - check status (2 values):
                fail (when the check is failed)
                ok   (when the check succeed)
timeout      - how much time to wait for the reply from the host until connection
               will be closed with the timeout. (value is in milliseconds)
fail_message - the reason why the check was failed
