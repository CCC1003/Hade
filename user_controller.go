package main

import "Hade/framework"

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}
