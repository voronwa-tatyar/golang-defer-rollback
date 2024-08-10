# Repository Explanation
This repository shows how to use defer as rollback pattern

This pattern would be used when `A component` is easy to set and rollback
And `B component` or `C component` has various failure cases

In `complexCase.go`, I used the case of gRPC or REST API response
that comes with status code in response

You can change the `setXComponent` function to compare the result of the functions

# Execution Result

## Simple Case
When setBComponent() returns nil
```
func setBComponent() error {
	return nil
}
```
```
Running simple case
Before execution, A is  0
After execution, A is  1

Before execution, A is  0
After execution, A is  1

Simple case complete
```

When setBComponent() returns error
```
Running simple case
Before execution, A is  0
failed to set b, reason :  sample error
After execution, A is  0

Before execution, A is  0
failed to set b, reason :  sample error
After execution, A is  0

Simple case complete
```

## Complex Case
When SetCComponent() returns Response{200, ""}, nil
```
func setCComponent() (Response, error) {
	return Response{200, ""}, nil
}
```
Running complex case
Before execution, A is  0
After execution, A is  1

Before execution, A is  0
After execution, A is  1

Complex case complete
```

When SetCComponent() returns Response{200, ""}, errors.New("sample error")
```
func setCComponent() (Response, error) {
	return Response{200, ""}, errors.New("sample error")
}
```
```
Running complex case
Before execution, A is  0
failed to set c, reason :  sample error
After execution, A is  0

Before execution, A is  0
failed to set c, reason :  sample error
After execution, A is  0

Complex case complete
```

When SetCComponent() returns Response{404, "not found error"}, nil
```
func setCComponent() (Response, error) {
	return Response{404, "not found error"}, nil
}
```

```
Running complex case
Before execution, A is  0
failed to set c, reason :  not found error
After execution, A is  0

Before execution, A is  0
failed to set c, reason :  not found error
After execution, A is  0

Complex case complete
```