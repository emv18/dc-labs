User Guide

1.- Installation to Set Up Environment
    a) Install Gin with the following command
    b) Write in terminal $ go get -u github.com/gin-gonic/gin
    c) Install jwt-go with the following command
    d) $ go get -u github.com/dgrijalva/jwt-go

2.- Start server
    a) Write $ go run server.go

3.- Login
    a) Write $ curl -u username:password localhost:8080/login
    b) Receive token

4.- Logout
    a) Write $ curl -H "Authorization: Bearer 'token'" http://localhost:8080/logout

5.- Upload
    a) Write $ curl -F 'data=@/path/to/local/image.png' -H "Authorization: Bearer 'token'" http://localhost:8080/upload

6.- Status
    a) Write $ curl -H "Authorization: Bearer 'token'" http://localhost:8080/status


==========
