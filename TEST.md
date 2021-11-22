# Testing ideas

~~~
echo "GET http://localhost:8080/" | vegeta attack -duration=5s | vegeta report
~~~


~~~
curl -c cookies.txt "http://127.0.0.1:8080/next"
curl -b cookies.txt "http://127.0.0.1:8080/next"
curl -b cookies.txt -c cookies.txt "http://127.0.0.1:8080/next" ; echo
~~~


~~~
ali -r 1000 http://127.0.0.1:8080/random
ali -r 5000 http://127.0.0.1:8080/random
echo "GET http://127.0.0.1:8080/random" | vegeta attack -rate=5000  -duration=10s | vegeta report
~~~

~~~
http_load -rate 1000 -seconds 10 url
(url file:http://127.0.0.1:8080/random)
~~~ 


