# http-rest-api-async-invocation-golang

This project shows an approach to expose a REST API which works asynchronously, in this project you can find:

1. Response with a UUID in a body, a consumer can do a long polling strategy in order to know the result of an execution
2. Response with UUID in a Location header, a consumer can do a long polling strategy in order to know the result of an execution
3. Request which allows you to check the status of any execution by using its UUID
4. Response which simulates an invocation to a URL callback

I left a postman collection, and a visual code project in order for you to test this API

If you have any doubt, do not hesitate in contacting me, my email is santiago.diaz5689@gmail.com
