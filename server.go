package main

import (
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

const classListData string = "classlist"

//The cache is just a string => bytes map, so all of the information we want to acces
//(what classes are available for querying, what levels of a certain subschool does a class have available to it, etc...)
//is stored behind a cacheKey.
//A cacheKey is simply the name of the entity we seek information on (wizardschools, wizardsubschools, clericevovationlevels, druidconjuration3spells)
type CacheKey string

//Holds all the information we need for the frontend to function
type server struct {
	//Data grab-bag where we keep all of our information relevant to spells.
	//This ranges from keys with empty values to tell us which classes exist,
	//to the list of all classes we know about, to which spells a particular class/school/subschool/level has
	magicCache map[string][]byte
	//Need a lock so that we can read/write to the cache safely
	magicCacheLock    sync.RWMutex
	classSchoolLevels map[string]map[string]map[string]int
	//spellsDB          string
	dbConnection *pgxpool.Pool
}

func newServer() (*server, error) {
	server := &server{
		magicCache:     make(map[string][]byte, 100),
		magicCacheLock: sync.RWMutex{},
	}
	if err := server.cacheClasses(); err != nil {
		return server, err
	}
	return server, nil
}

func (c *server) isValidClass(thisClass string) bool {
	_, class := c.magicCache[thisClass]
	return class
}

func (c *server) isValidSchool(class, cacheKey string) bool {
	if c.isValidClass(class) {
		if _, validSchool := c.magicCache[cacheKey]; validSchool {
			return true
		}
		return false
	}
	return false
}
