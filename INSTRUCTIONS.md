# Description
This is the first deliverable for the golang bootcamp 2020. This project uses the restcountries API, we use it to retrieve info 
about countries (name, capital, population, etc).

# How to run
Set the following env variable with an valid directory. This is where CSV files will be stored.
```sh
export COUNTRIES_FILE="..."
```

To run the project: 
```go
  go run main.go
```
The server will run in localhost:8000.
# Dependencies
Third projects uses 
- Gin HTTP Framework https://github.com/gin-gonic/gin
- Go Convey (testing framework) https://github.com/smartystreets/goconvey

# Endpoints and how to use
There is two main endpoints:
```
GET localhost:8000/v1/:country_code/fetch <- Fetch from API and store the data file. 
GET localhost:8000/v1/:country_code <- Read the CSV file. 
```
Both endpoints responds with JSON response. 
The country code is ISO-1136 format (https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes).
For example 
```
GET localhost:8000/v1/countries/mx/fetch
GET localhost:8000/v1/countries/mx
```

NOTE: You do need to trigger the fetch endpoint before reading the csv the first time, so the file with the country info is stored. Otherwise 
you'll get an 500 Internal Server Error code because there is not file with that name (i.e mx.csv)

# Response 
This is the struct of the response: 
```json
{
    "Name": "Mexico",
    "Capital": "Mexico City",
    "Area": 1964375,
    "Gini": 47,
    "NativeName": "México",
    "NumericCode": "484",
    "Population": 122273473,
    "Region": "Americas",
    "SubRegion": "Central America",
    "Code": "mx"
}
```

# Testing
To run the tests just go to 
```sh
/handlers/v1/countries/handlers
/models
```

and run 
```go
    go test
```

# TO DO: 
- Handle HTTP errors better.
- Add better HTTP request testing. 
- More documentation for exported functions.