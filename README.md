
# NSE Board Meetings

A Simple GOLang application using ReactJS and data from NSE corporate board meeting.

## Structure

* `frontend`: ReactJS App to display data (currently uses v2 api).
* `v1nse`: Legacy NSE website (www1) due for End of Life in August 2020. Used `go-colly` for scraping data of the website.
* `v2nse`: Used JSON APIs from nse website to display meaningful data. (Selective JSON data is received rather than entire JSON).
* `main.go`: Uses `gin` router to map v1 and v2 apis accurately for the required purposes.

## Running Locally

### Backend

Make sure you have [Go](http://golang.org/doc/install) version 1.12 or newer. Preferrably OS is **Mac OS**

```sh
$ git clone https://github.com/hjoshi123/NSEBoardMeetings.git
$ cd NSEBoardMeetings
$ export PORT=5000
$ go run main.go
[GIN-debug] GET   /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (2 handlers)
[GIN-debug] HEAD  /static/*filepath         --> github.com/gin-gonic/gin.(*RouterGroup).createStaticHandler.func1 (2 handlers)
[GIN-debug] GET   /                         --> main.main.func1 (2 handlers)
[GIN-debug] POST  /getBoardMeetings         --> main.getBoardMeetings (2 handlers)
[GIN-debug] Listening and serving HTTP on :5000
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

### Frontend

Before running the frontend, make sure the backend is up and running at PORT 5000 (by default). If you use a different port, change it accordingly in the react app.

```sh
$ cd frontend
$ yarn
$ yarn start
Compiled successfully!

You can now view frontend in the browser.

  Local:            http://localhost:3000
  On Your Network:  http://192.168.1.1:3000

Note that the development build is not optimized.
To create a production build, use yarn build.
```

## Running through Docker

To run through docker, ensure docker daemon is running. Once the docker daemon is up, execute the following command in terminal

```sh
$ docker-compose up -d --build
Building backend
Step 1/11 : FROM golang:latest AS builder
 ---> 75605a415539
Step 2/11 : ADD . /app
 ---> ff4ff21b9bc8
Step 3/11 : WORKDIR /app
 ---> Running in 42d4b241f7fd
Removing intermediate container 42d4b241f7fd
 ---> 014ed3b6baf2
Step 4/11 : RUN go mod download
 ---> Running in 38c0dcf955fb
...
...
Step 12/13 : EXPOSE 80
 ---> Using cache
 ---> 6de540b07ded
Step 13/13 : CMD ["nginx", "-g", "daemon off;"]
 ---> Using cache
 ---> 8622c0202783
Successfully built 8622c0202783
```

The frontend is mapped to port 80. So once the docker-compose runs successfully, [localhost](http://localhost) will contain your app.
