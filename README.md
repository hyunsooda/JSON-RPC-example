# JSON-RPC-example
<hr>
This example is shown how to serve RPC with JSON

There is an embeding C function to provide user information in package folder(userinfo)


<pre><code>
SERVER : go run server.go 8080
CLIENT(curl) : curl -X GET -H "Content-Type: application/json" -d '{"method": "Testlib.GetInfo", "params": [] }' http://localhost:8080/info
CLIENT(http post request) : http://localhost:8080/info 
</code></pre>

If you choice second way in client, you can't experiment in browser directly 
because typing in url input box is just to request GET method 
So you need a popular tool such as https://apitester.com/ and do set header to application/json and empty parameter



