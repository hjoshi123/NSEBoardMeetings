
# NSE Board Meetings

A Simple GOLang application using HTML and data from NSE corporate board meeting.

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.12 or newer

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

## Deployement