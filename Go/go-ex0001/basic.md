# Developing a RESTful API with Go and Gin

## GO-EX0001

Table of Contents
Prerequisites
Design API endpoints
Create a folder for your code
Create the data
Write a handler to return all items
Write a handler to add a new item
Write a handler to return a specific item
Conclusion
Completed code
This tutorial introduces the basics of writing a RESTful web service API with Go and the Gin Web Framework (Gin).

You’ll get the most out of this tutorial if you have a basic familiarity with Go and its tooling. If this is your first exposure to Go, please see Tutorial: Get started with Go for a quick introduction.

Gin simplifies many coding tasks associated with building web applications, including web services. In this tutorial, you’ll use Gin to route requests, retrieve request details, and marshal JSON for responses.

The tutorial includes the following sections:

Design API endpoints.
Create a folder for your code.
Create the data.
Write a handler to return all items.
Write a handler to add a new item.
Write a handler to return a specific item.

Good luck..

### Linkedin: <a href="https://www.linkedin.com/in/roham-shahedi">ROHAM SHAHEDI</a>

## Github: <a href="https://www.linkedin.com/in/roham-shahedi">ROHAM96</a>

Run:

```
go mod init main
go mod tidy
go run .
```

Now! your basic web server started On:
<a href="http://localhost:8000">http://localhost:8000</a>
#
### 1- To Add Article

```
curl --location --request POST 'http://localhost:8000/articles' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "4",
    "title": "White Train",
    "body": "White Train body string text",
    "author": "4"
}'
```
Result:
```
{
    "id": "4",
    "title": "White Train",
    "body": "White Train body string text",
    "author": "4"
}
```

### 2- To Get Article By ID
```
curl --location --request GET 'http://localhost:8000/articles/2'
```
Result:
```
{
    "id": "2",
    "title": "Green Train",
    "body": "Green Train body string text",
    "author": "2"
}
```
### 3- To Get All Articles
```
curl --location --request GET 'http://localhost:8000/articles'
```
Result:
```
[
     {
        "id": "1",
        "title": "Blue Train",
        "body": "Blue Train body string text",
        "author": "1"
    },
    {
        "id": "2",
        "title": "Green Train",
        "body": "Green Train body string text",
        "author": "2"
    },
    {
        "id": "3",
        "title": "Red Train",
        "body": "Red Train body string text",
        "author": "3"
    },
    {
        "id": "4",
        "title": "White Train",
        "body": "White Train body string text",
        "author": "4"
    }
]
```
#
is typing..|
