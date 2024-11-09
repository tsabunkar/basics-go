package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string { // Hello <- func name, (name string) <- argument and its datatype string , ) string { <- complete func datatype 
    /* 
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name) // instead of separately declaraing the variable and intilazing it use := operator
 */
    message := fmt.Sprintf("Hi, %v. Welcome!", name) // %v format verb
    return message
}
