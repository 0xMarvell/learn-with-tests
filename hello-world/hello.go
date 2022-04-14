// Package hello sends a simple greeting to the user.
package hello

const prefix = "Hello, "

//Hello returns a simple greeting using the name inputted by the user.
func Hello(name string) string {
	if name == "" {
		return "Hello"
	}
	return prefix + name
}
