# cmpe273-Assignment2

This assignment contains RESTful APIs in Golang with the basic CRUD operations. It is a address directory
where you can add, update, get and delete the address.

The webservice fetches the co-ordinates data from Google Maps API and saves to the Mongo Labs.

Following are the sample Request and Responses:

POST: /locations/
Request:

{
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113"
}

Response:
{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113",
   "coordinate" : { 
      "lat" : 38.4220352,
     "lng" : -222.0841244
   }
}

Get: /locations/{location_id}
Response:
{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113",
   "coordinate" : { 
      "lat" : 38.4220352,
     "lng" : -222.0841244
   }
}

Update: /locations/{location_id}
Request:
{
   "address" : "1600 Amphitheatre Parkway",
   "city" : "Mountain View",
   "state" : "CA",
   "zip" : "94043"
}

Response:
{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "1600 Amphitheatre Parkway",
   "city" : "Mountain View",
   "state" : "CA",
   "zip" : "94043"
   "coordinate" : { 
      "lat" : 37.4220352,
     "lng" : -122.0841244
   }
}

Delete: /locations/{location_id}
