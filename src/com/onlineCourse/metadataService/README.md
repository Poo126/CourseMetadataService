# CourseMetadataService

This is simple example **gorilla/mux**.  

## Install and Run
```shell
$ go get github.com/kkamdooong/go-restful-api-example

$ cd $GOPATH/src/github.com/Poo126/src/com/onlineCourse/metadataService
$ go build
$ ./metadataService
```

## API Endpoint
- http://localhost:3000/api/v1/courses
    - `GET`: get list of courses
    - `POST`: create course
- http://localhost:3000/api/v1/courses/{name}
    - `GET`: get course
    - `PUT`: update course
    - `DELETE`: remove course

## PUT example
curl --request POST --url http://localhost:3000/api/v1/courses --header 'content-type: application/json' --data 
'{ 
  "name": "golang",
  "authorId": "suveetha",
"version": "1", "state": "provisioned"}'