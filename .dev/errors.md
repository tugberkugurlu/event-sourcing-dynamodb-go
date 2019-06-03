You see this when the table already exists in dynamodb:

```
➜  example git:(master) ✗ go run main.go
2019/06/03 08:19:05 ResourceInUseException: Cannot create preexisting table
	status code: 400, request id: 35c2d8ac-d749-4903-8266-0f0f4d72e353
exit status 1
```