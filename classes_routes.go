package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (s *server) classesList(c echo.Context) error {
	return c.Blob(200, "Application/json", s.magicCache[classListData])
}

func (s *server) classSchoolList(c echo.Context) error {
	className := c.Param("name")
	if !s.isValidClass(className) {
		return fmt.Errorf("%s is not a valid class name", className)
	}
	retList := make([]string, 0, len(s.classSchoolLevels[className]))
	for schoolName := range s.classSchoolLevels[className] {
		retList = append(retList, schoolName)
	}
	return c.JSON(200, retList)
}

func (s *server) classSchoolSubschoolList(c echo.Context) error {
	className := c.Param("name")
	schoolName := c.Param("school")
	if !s.isValidSchool(className, schoolName) {
		return fmt.Errorf("%s is not a valid magic school for class %s", schoolName, className)
	}
	return nil
}
