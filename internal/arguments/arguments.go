package arguments

import "github.com/WAY29/hackbox/internal/tools"

var globalArgumentMap = map[string]*tools.Arg{}

func Save(value *tools.Arg) {
	globalArgumentMap[value.Name] = value
}

func Remove(name string) bool {
	if _, ok := globalArgumentMap[name]; ok {
		delete(globalArgumentMap, name)
		return true
	}
	return false
}

func Get(name string) *tools.Arg {
	if value, ok := globalArgumentMap[name]; ok {
		return value
	}
	return nil
}
