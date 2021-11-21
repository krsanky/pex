# Web Backend Technical Challenge

Please design and implement a web based API that steps through the Fibonacci sequence. 

The API must expose 3 endpoints that can be called via HTTP requests:
* current - returns the current number in the sequence
* next - returns the next number in the sequence
* previous - returns the previous number in the sequence

Example:
```
current -> 0
next -> 1
next -> 1
next -> 2
previous -> 1
```

### Requirements:

* The API must be able to handle high throughput (~1k requests per second).
* The API should also be able to recover and restart if it unexpectedly crashes.
* Assume that the API will be running on a small machine with 1 CPU and 512MB of RAM.
* You may use any programming language/framework of your choice.

---

## Assumptions:

1. next increments, current ... does previous decrement it?
2. is there common state, or state per client?
1. State is per client.
	Otherwise multiple clients are racing each other, and is the req/resp
	cycle atomic to handle that.
	If there is common state, then wrap the current -> next transition in a semaphore.

## Links
- <https://github.com/tsenart/vegeta>
- <https://github.com/alexedwards/scs>

- <https://en.wikibooks.org/wiki/Algorithm_Implementation/Mathematics/Fibonacci_Number_Program>
- <https://medium.com/future-vision/fibonacci-sequence-algorithm-5eebae4e85be>


## Testing

~~~
echo "GET http://localhost:8080/" | vegeta attack -duration=5s | vegeta report
~~~

---

~~~
curl -c cookies.txt "http://127.0.0.1:8080/next"
curl -b cookies.txt "http://127.0.0.1:8080/next"
curl -b cookies.txt -c cookies.txt "http://127.0.0.1:8080/next" ; echo
~~~

---

## More Testing ideas

~~~
ali -r 1000 http://127.0.0.1:8080/random
ali -r 5000 http://127.0.0.1:8080/random
echo "GET http://127.0.0.1:8080/random" | vegeta attack -rate=5000  -duration=10s | vegeta report
~~~

http_load -rate 1000 -seconds 10 url
(url file:http://127.0.0.1:8080/random)

