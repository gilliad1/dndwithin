package main

import (
	"github.com/labstack/echo/v4"
)

func (c *server) GetEchoServer() (*echo.Echo, error) {
	e := echo.New()
	e.GET("/classes/list", c.classesList)
	e.GET("/classes/:name/magic/schools", c.classSchoolList)
	e.GET("/classes/:name/magic/:school/subschools", c.classSchoolSubschoolList)
	return e, nil
}

func (c *server) cacheClasses() error {
	c.magicCacheLock.Lock()
	defer c.magicCacheLock.Unlock()
	for _, class := range []string{""} {
		c.magicCache[class] = []byte{}
	}
	return nil
}
