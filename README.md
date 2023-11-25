# Test_FizzBuzz

## Prerequisite

The go verison used is 1.18  
The linter used is golangci-lint with version 1.51.0

## Build

```bash
go build .
```

Will compile the microservice.

## Run

```bash
go run .
```

Will execute without compiling the microservice.

## Production Ready

The application is ready for production, it has been well tested, is considered stable and reliable enough for high capacity usage.

## API FizzBuzz

This API prints the numbers from 1 to 200 by default, but for multiples of seven print "Fizz" instead of the number and for the multiples of nine print "Buzz".
For numbers which are multiples of both seven and nine print "FizzBuzz".
Most of those parameters a modifiable.

For more information about the API please check the swagger file `openapi.yaml`

## Useful

Please visit the `Makefile` for some quick command
