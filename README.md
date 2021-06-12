# bdj-muhammadnurbasari
**this application service for test on Jakarta Smart City**

**How to run project** 

1. create database "bdj-muhammadnurbasari" on mysql
2. fill value in files config/sample.env ( default on : localhost:12345 )
3. run command "go mod vendor"
4. run migration table using command "go run database/main.go --up"
5. run service using command "go run main.go"
6. Documentation api {baseURL}/docs/swagger/index.html


**How to Hit Endpoint**

API untuk get list Combine Data Rs & Data kelurahan

**URL** : `/bdj/data/combine`

**Method** : `GET`

## Request 
**Header**
**Content-Type** : `application/json`<br>
**Authorization** : `$2a$14$8kYmqb44a9UpRocodQUSi.sBPKSRRPFBMYuKnFHlNqJtnXbLi0p3C`<br>



You can Try on Documentasi swagger or Using Posmant

use 
**Authorization** : `$2a$14$8kYmqb44a9UpRocodQUSi.sBPKSRRPFBMYuKnFHlNqJtnXbLi0p3C`