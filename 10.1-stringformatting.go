package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value interface{} // The value can be of any type
	isSet bool
}

/*
%s - for strings
%d - for integers
%f - for floating-point numbers, %.2f for 2 decimal places
%t - for booleans
%T - for the type of a value
%v - for any value in a default format
%+v - for any value with field names (useful for structs)
%#v - for any value in Go syntax representation
%% - to print a literal percent sign
*/

// Implementing the Stringer interface for ConfigItem struct
// Where you have to return a string use Sprintf
// since it returns a string instead of printing it directly as Printf does.
func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s\n Value: %v%%\n IsSet: %t\n", c.Key, c.Value, c.isSet)
}

func main() {

	c := ConfigItem{
		Key:   "CPUUsageInPercent",
		Value: 67.5,
		isSet: true,
	}

	fmt.Println(c) // Calls the String() method to get the string representation of the ConfigItem struct

}
