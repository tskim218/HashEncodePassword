# HashEncodePassword

Hash and Encode a Password String
1. handler /hash endpoint with a form field named password to provide
the value to hash. An incrementing identifier will return immediately but the password will
not be hashed for 5 seconds. The hash will be computed as base64 encoded string of the
SHA512 hash of the provided password.
For example, the first request to:
curl —data “password=angryMonkey” http://localhost:8080/hash
will return 1 immediately. The 42nd request will return 42 immediately.

2. handelr /hash/identifire will return the encoded password. 5 seconds after the POST to /hash that returned
an identifer you should be able to curl
http://localhost:8080/hash/42 and get back the value of
“ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZP
ZklJz0Fd7su2A+gf7Q==”

Statistics End-Point
1. handler /stats provide a statistics endpoint to get basic information about your password hashes.
A GET request to /stats will return a JSON object with 2 key/value pairs. The “total”
key will have a value for the count of POST requests to the /hash endpoint made to the
server so far. The “average” key will have a value for the average time it has taken to
process all of those requests in microseconds.
For example: curl http://localhost:8080/stats should return something like:
{“total”: 1, “average”: 123}

Graceful Shutdown
1. handler /shutdown provides support for a “graceful shutdown request”. If a request is made to /shutdown the
server will reject new requests. The program will wait for any pending/in-flight work to
finish before exiting.

In Memory DB and Worker
1. In order to keep tracking of the requests, in-memory db and worker structs are
used.

Instructions.
1. Start the server. $> go run main.go
2. Open another terminal to request hash and encode the password
$> curl --data "password=hello" http://localhost:8080/hash
3. Open another terminal or use the same to request encode password by identifier
$> curl http://localhost:8080/1
4. Open another terminal or use the same to request statistics
$> curl http://localhost:8080/stats
5. Open another terminal or use the same to request shut down gracefully
$> curl http://localhost:8080/shutdown

Simulating a graceful shut down by following steps.
1. Comment in line 52 - 55 in get_password.go
2. Start the sever $> go run main.go
3. Follow the instruction 2 and 3 above
4. Run shutdown handler asap $> curl http://localhost:8080/shutdown
You should be able to see the server is waiting until current handers are done.
Also any new requests will be rejected.
