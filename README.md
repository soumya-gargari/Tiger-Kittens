# Tiger-Kittens

1....
# End point for creating a user : http://localhost:8080/user (POST request)
# Request Body :
==============================

{
  "username": "soumya",
  "password": "sdponV2X@8961816991",
  "email": "soumyagargari1@gmail.com"
 }

# Response Body:
==============================
 if successfully created, it will get 201 created response else will get 500 internal server error on internal failure
 or 400 bad request for passing wrong data

2....
# End point for login of a user : http://localhost:8080/login (GET request)
# Query Parameters:
==============================

  "username": "soumya",
  "password": "sdponV2X@8961816991",

# Response Body:
==============================
 if successfully login, it will give 200 ok response else will get 500 internal server error on internal failure
 or 400 bad request for passing wrong query data along with actual error message

3....
# End point for creating a tiger : http://localhost:8080/tiger (POST request)
# Request Body:
==============================

 {
  "name": "tiger20",
  "dateOfBirth": "30/12/1994",
  "lastSeen": "2019-09-05T22:16:20Z",
  "lastSeenlatitude": "13.90",
  "lastSeenlongitude": "10.60"
}

# Response Body:
==============================
 if successfully created, it will give 201 created response else will get 500 internal server error on internal failure
 or 400 bad request for passing wrong data along with actual error message

4....
# End point for getting all tigers : http://localhost:8080/tiger (GET Request)

# Response Body:
==============================
 [
  {
    "name": "tiger10",
    "dateOfBirth": "",
    "lastSeen": "2019-09-05T22:16:20Z",
    "lastSeenlatitude": "12.90",
    "lastSeenlongitude": "9.60"
  },
  {
    "name": "tiger20",
    "dateOfBirth": "30/12/1994",
    "lastSeen": "2019-09-05T22:16:20Z",
    "lastSeenlatitude": "13.90",
    "lastSeenlongitude": "10.60"
  }
]

5....
# End point for creating a tiger sighting data : http://localhost:8080/tiger/sighting (POST request)
# Request Body:
==============================

 {
  "name": "tiger20",
  "timestamp": "2024-01-12T15:30:00Z",
  "latitude": "14.97",
  "longitude": "15.59",
  "uploadImage": true
}

# Response Body:
==============================
 if successfully created, it will give 201 created response else will get 500 internal server error on internal failure
 or 400 bad request for passing wrong data along with actual error message



 6....
# End point for getting all tigers sighting data : http://localhost:8080/tigers/sighting (GET Request)

# Response Body:
==============================
[
  {
    "name": "tiger10",
    "timestamp": "2024-01-12T15:30:00Z",
    "latitude": "20.97",
    "longitude": "8.59",
    "uploadImage": true
  },
  {
    "name": "tiger20",
    "timestamp": "2024-01-12T15:30:00Z",
    "latitude": "14.97",
    "longitude": "15.59",
    "uploadImage": true
  }
]