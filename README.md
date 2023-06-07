# The Task
This is a production ready web service which combines two existing web services.  
Fetching a random name from https://names.mcquay.me/api/v0
Fetching a random Chuck Norris joke from http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=John&lastName=Doe  
Combines the results and returns them to the user.

# How to run
Requires go to be installed locally
## Starting the webserver
Navigate to the project directory from the command line. 
Webserver can be started either using go run:
```
go run main.go name.go joke.go
```
Or build the executable:
```
go build main.go name.go joke.go
```
After the command completes, run the executable file that is produced

## Using the new web service
The web service will return a value via curl, or opening http://localhost:5000 in your preferred browser.
```
$ curl "http://localhost:5000"
Hasina Tanweer’s OSI network model has only one layer - Physical..
```

# Tests
Basic unit tests are included. API endpoints were not mocked for these tests, but improvements could be made by mocking both the name and joke http requests. Functional tests would be nice to add, but this project is relatively simple.  
The tests can be run from the command line once you have navigated to the project directory:
```
go test
```

# Implementation notes
## Load testing
Load testing was performed with apache bench, but the name endpoint began to return the string "You have reached maximum request limit." instead of the expected JSON response. Further development should check the response type and handle the error more appropriately if either of the APIs returns an unexpected response.

## Time spent
The code itself (including the unit tests) took about 3 hours. Load testing and README updating took some additional time.

## Note on concurrency:
The task stated that the webserver should handle concurrent connections, which I expected to mean I would need to fire off explicit go routines, however, research into http.ListenAndServe shows that Serve is a handled in a go routine so this was not necessary.
