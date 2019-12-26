#Getting Started

## the backend

the backend for this excercise is a golang app. The resources for launching the backend are located in pinkaloo_backend/src/github.com/gitJucy/appServer. Build at your GOPATH and run the executable 'appserver'. This will open the service at port 8080. There is a Postman collection JSON at the root of this project that will allow you access the endpoints in this golang service. The Golang servicve persists app metadata, validates metadata when adding, and allows simple keyword searches. Be sure to add App data in Postman before using the frontend.

## the frontend

The frontend is a react application that provides a _very_ simple UI to explore the search function of the webservice. This service can be launched in a dev state by running npm start from pinkaloo_frontend.
