package middleware

import (
	"Hade/framework"
	"fmt"
)

func Test1() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}
func Test2() framework.ControllerHandler {
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}
