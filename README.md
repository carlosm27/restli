# Restli
A REST client CLI tool. Built it with Cobra and Resty

 ### GET
 ``
 restli get --url <url>
 ``
 
 ### POST
 
 ``
 restli post --url <url> --body <body>
 ``
Windows users must add a "\" before every quote in the body, otherwise, it won't work and it will returns a Bad Request status code.

Example:

``
restli post --url "http://localhost:8000/api" --body {\"name\":\"Carlos\", \"lastname\":\"Marcano\"}
``
