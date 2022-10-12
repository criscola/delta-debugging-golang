# Delta debugging (1-minimal) 

Simple delta debugging (1-minimal variant) example written in Go during the S&DE Atelier: Software Analytics course at Università della Svizzera italiana.

Credits to Prof. Dr. Andreas Zeller for inventing this clever technique.

Running with Go installed:

```bash
go run main.go htmlPage.txt failForHtml.txt
```

Running in Docker:

```bash
docker run --rm -v $(pwd):/delta golang:latest go run /delta/main.go /delta/htmlPage.txt /delta/failForHtml.txt
```
