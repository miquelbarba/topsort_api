Calculate API
=============
To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried.

The API should accept a request that includes a list of flights, which are defined by a source and destination airport code.

These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

Run Server
----------
Requires Goland 1.19.

To start the server in port 8080:

    $ go run cmd/main.go

Usage
-----

The server implements only one endpoint /calculate that is only accessible via the HTTP GET method.

This endpoint has one parameter "path" received as a query param. This parameter is an array, where each item is one flight with the origin and destination separated with a comma. 

When successful, an HTTP 200 status is returned with a JSON in the body that contains an array with two items, the first is the origin airport and the second item is the final airport.

Example:

    $ curl http://localhost:8080/calculate?path=IND,EWR&path=SFO,ATL&path=GSO,IND&path=ATL,GSO

    {"result":["SFO","EWR"]}

In the case of an error, the HTTP status indicates the error and in the body the message of the error

    $ curl http://localhost:8080/calculate?path=IND,EWR&path=SFO,ATL&path=GSO,IND&path=ATL

    {"message":"invalid parameter ATL"}

Tests
-----

To run the tests:

    $ go test ./...
