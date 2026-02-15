package main

// VARIABLE NAMING RULES

func main() {
	// CORRECT DECLARATIONS
	var speed int
	var SpeeD int

	// underscore is allowed but not recommended
	var _speed int

	// Unicode letters are allowed
	var kocak int

	// keep the compiler happy
	_, _, _, _ = speed, SpeeD, _speed, kocak
}
