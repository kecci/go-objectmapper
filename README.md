# go-objectmapper

Object mapping to another object with devfeel/mapper

Usage:
```go
	userData := &UserData{ // Have 2 Field
		ID:   1,
		Name: "John",
	}
	userResponse := &UserResponse{} // Have 3 Field
	
    // Map userData -> userResponse
	err := mapper.Mapper(userData, userResponse)
```

Output:
```sh
2021/04/20 02:36:47 {"ID":1,"Name":"John","Note":""}
2021/04/20 02:36:47 {"ID":2,"Name":"Doe"}
```

Source: https://github.com/devfeel/mapper