package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func abac() {
	e, _ := casbin.NewEnforcer("../config/model.conf", "../config/policy.csv")

	// Print the policy to log.
	e.EnableLog(true)

	// Check the permission.
	sub := "alice" // the user that wants to access a resource.
	obj := "data2" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	res, err := e.Enforce(sub, obj, act)
	if err != nil {
		panic(err)
	}

	if res {
		// permit alice to read data1
		fmt.Println("permit alice to read data1")
	} else {
		// deny the request, show an error
		fmt.Println("deny the request, show an error")
	}
}
