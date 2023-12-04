package middleware

import (
	"Hade/Hin"
	"fmt"
)

func Test1() Hin.ControllerHandler {
	return func(c *Hin.Context) error {
		fmt.Println("middleware pre test1")
		c.Next()
		fmt.Println("middleware post test1")
		return nil
	}
}
func Test2() Hin.ControllerHandler {
	return func(c *Hin.Context) error {
		fmt.Println("middleware pre test2")
		c.Next()
		fmt.Println("middleware post test2")
		return nil
	}
}
