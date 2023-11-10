# Go_Microservice_with_CI-CD
Go Microservice backend application with tests

Uses a static data instead of backend, on running the application, 

- http://localhost:8080/ return HTML page
- http://localhost:8080/article/view/<id> where id is 1/2/ returns the article data
- curl -X GET -H "Accept: application/json" http://localhost:8080/ returns articles in JSON format
- curl -X GET -H "Accept: application/xml" http://localhost:8080/article/view/1 returns a individual article in XML format

References:-
 - https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin#h-displaying-a-single-article
