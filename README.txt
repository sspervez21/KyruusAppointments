This is an implementation that uses go-swagger to generate server code using the definition provided in the file swagger.xml.

While most of the server code is auto-generated my code mainly lives under the app\ folder.

I have chosen to use an in memory implementation of the backend for simplicity and ease of execution.

Known issues:
  - Update doctor: This action is under specified. Today we simply overwrite the existing information with the new information. Recommended that we implement an algorightm that merges the existing data with the new data.
  
* What are some real-world constraints to booking appointments that would add complexity to this API and how would they impact the design.

One of the major concerns is multiple users trying to access the same data at the same time and how we should address any overlapping requests. This is tricky because if two users try to book conflicting appointments we will need a way to allow this behavior (or gracefully disallow).
A possible solution would be to allow users to enter multiple possible appointments at the same time and pick (perhaps in order of preference) the one that works.

Requirements:
Go : https://golang.org/dl/
Go Swagger : go get -u -v github.com/go-swagger/go-swagger/cmd/swagger
dep : https://github.com/golang/dep/blob/master/README.md

### generate swagger files
swagger generate server -f swagger.yml

### download dependences
dep init
dep ensure

### build the server
go build ./cmd/kyruus-server/main.go

### start the server
./main.exe --port=8082

### sample usage
Create Doctor:
curl -X POST http://localhost:8082/doctors -H "accept: application/json" -H "Content-Type: application/json" -d @input\Doctor1.json
curl -X POST http://localhost:8082/doctors -H "accept: application/json" -H "Content-Type: application/json" -d @input\Doctor2.json

Get Doctor:
curl -X GET http://localhost:8082/doctors/1

Update Doctor:
curl -X PUT http://localhost:8082/doctors/1 -H "accept: application/json" -H "Content-Type: application/json" -d @input/Doctor2.json

Delete Doctor:
curl -X DELETE http://localhost:8082/doctors/1

Create Appointment:
curl -X POST http://localhost:8082/appointments -H "accept: application/json" -H "Content-Type: application/json" -d @input\Appointment1.json
curl -X POST http://localhost:8082/appointments -H "accept: application/json" -H "Content-Type: application/json" -d @input\Appointment2.json
curl -X POST http://localhost:8082/appointments -H "accept: application/json" -H "Content-Type: application/json" -d @input\Appointment3.json
curl -X POST http://localhost:8082/appointments -H "accept: application/json" -H "Content-Type: application/json" -d @input\Appointment4.json

Get Appointments for a doctor with this ID:
curl -X GET http://localhost:8082/appointments/1

Delete Appointment:
curl -X DELETE http://localhost:8082/appointments -H "accept: application/json" -H "Content-Type: application/json" -d @input\Appointment2.json

