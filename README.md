To test, run `make`. The file `main_test.go` contains type assertions that demonstrate the bug.

## goa v1.1.0

```
---> rm -rf main.go greeting.go extended_greeting.go app/ client/ swagger/ tool/
---> goagen bootstrap --design=github.com/xoob/goa-testcase-payloads/design
...
---> go test main_test.go
# command-line-arguments_test
./main_test.go:22: cannot use (&app.ExtendedGreeting literal).ID (type string) as type int in assignment
./main_test.go:42: cannot use (&app.ExtendedGreetingMedia literal).ID (type string) as type int in assignment
./main_test.go:63: cannot use (&app.CreateExtendedGreetingPayload literal).ParentOptional (type *string) as type *bool in assignment
./main_test.go:72: cannot use (&app.CreateOptionalExtendedGreetingPayload literal).ParentOptional (type *string) as type *bool in assignment
FAIL	command-line-arguments [build failed]
---> echo 'static tests failed'
static tests failed
```

## goa v1.2.0

```
$ make
---> rm -rf main.go greeting.go extended_greeting.go app/ client/ swagger/ tool/
---> goagen bootstrap --design=github.com/xoob/goa-testcase-payloads/design
...
---> go test main_test.go
# command-line-arguments_test
./main_test.go:57: cannot use (&app.CreateGreetingPayload literal).ParentOptional (type *string) as type *bool in assignment
./main_test.go:63: cannot use (&app.CreateExtendedGreetingPayload literal).ParentOptional (type *string) as type *bool in assignment
./main_test.go:65: cannot use (&app.CreateExtendedGreetingPayload literal).Active (type string) as type bool in assignment
./main_test.go:66: cannot use (&app.CreateExtendedGreetingPayload literal).Optional (type *string) as type *int in assignment
./main_test.go:72: cannot use (&app.CreateOptionalExtendedGreetingPayload literal).ParentOptional (type *string) as type *bool in assignment
./main_test.go:74: cannot use (&app.CreateOptionalExtendedGreetingPayload literal).Active (type string) as type bool in assignment
./main_test.go:75: cannot use (&app.CreateOptionalExtendedGreetingPayload literal).Optional (type *string) as type *int in assignment
FAIL	command-line-arguments [build failed]
---> echo 'static tests failed'
static tests failed
```
