# Restli
A REST client CLI tool. Built it with Cobra and Resty

 ### GET
 ``
 restli GET --url <url>
 ``
 
 ### POST
 
 ``
 restli POST --url <url> --body <body>
 ``
Windows users must add a "\" before every quote in the body, otherwise, it won't work and it will returns a Bad Request status code.

Example:

``
restli POST --url "http://localhost:8000/api" --body "{\"name\":\"Carlos\", \"lastname\":\"Marcano\"}"
``

### PUT

 ``
 restli PUT --url <url> --body <body>
 ``
 
 Windows users must add a "\" before every quote in the body, otherwise, it won't work and it will returns a Bad Request status code.

Example:

``
restli PUT --url "http://localhost:8000/api" --body "{\"name\":\"Carlos\", \"lastname\":\"Marcano\"}"

### PATCH

 ``
 restli PATCH --url <url> --body <body>
 ``
 
 Windows users must add a "\" before every quote in the body, otherwise, it won't work and it will returns a Bad Request status code.

Example:

``
restli PATCH --url "http://localhost:8000/api" --body "{\"name\":\"Carlos\", \"lastname\":\"Marcano\"}"

### DELETE

 restli DELETE --url <url>


