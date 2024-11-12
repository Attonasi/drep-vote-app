package validators

import (
	"example/hello/entity"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateVote(field validator.FieldLevel) bool {
	// Get the parent structure (Vote in this case)
	parent := field.Parent()

	// Find the Author field in the parent structure
	authorField, _ := parent.Type().FieldByName("Author")
	// Reflect to get the Person struct
	personValue := reflect.ValueOf(parent.Interface()).FieldByIndex(authorField.Index)

	// Get the Age field from Person
	ageField, _ := personValue.Type().FieldByName("Age")
	ageValue := personValue.FieldByIndex(ageField.Index)

	// Print the age
	fmt.Printf("Age: %d\n", ageValue.Int())

	// Proceed with your validation logic
	if person, ok := field.Parent().FieldByName("Author").Interface().(entity.Person); ok {
		return strings.Contains(field.Field().String(), GetValidatorString(uint8(person.Age)))
	}
	return false // If we can't get the age, fail validation
}

func GetValidatorString(valIndex uint8) string {
	var validatorStrings [11]string

	validatorStrings[0] = "zero"
	validatorStrings[1] = "one"
	validatorStrings[2] = "two"
	validatorStrings[3] = "three"
	validatorStrings[4] = "four"
	validatorStrings[5] = "five"
	validatorStrings[6] = "six"
	validatorStrings[7] = "seven"
	validatorStrings[8] = "eight"
	validatorStrings[9] = "nine"
	validatorStrings[10] = "ten"

	return validatorStrings[valIndex]
}
