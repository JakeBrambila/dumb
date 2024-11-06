package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Define custom types for HairColor and EyeColor
type HairColor string
type EyeColor string

// Define HairColor constants
const (
	Brown  HairColor = "brown"
	Blonde HairColor = "blonde"
	Black  HairColor = "black"
	Red    HairColor = "red"
)

// Define EyeColor constants
const (
	BrownEye EyeColor = "brown"
	Blue     EyeColor = "blue"
	Green    EyeColor = "green"
	Hazel    EyeColor = "hazel"
	BlackEye EyeColor = "black"
)

// Girl represents an individual with certain characteristics that affect difficulty scoring.
type Girl struct {
	name      string
	height    int
	weight    int
	age       int
	hairColor HairColor
	eyeColor  EyeColor
	college   bool
	sorority  bool
}

// determineDifficulty calculates a difficulty score based on the attributes of a Girl struct.
//
// Parameters:
//   - girl: A Girl struct representing the individual's characteristics.
//
// Returns:
//   - int: A difficulty score based on specific conditions related to the Girl struct's fields.
func determineDifficulty(girl Girl) int {
	difficulty := 0
	// Increase difficulty based on name characteristics
	if strings.HasPrefix(girl.name, "A") || strings.HasSuffix(girl.name, "a") {
		difficulty += 1
	}
	// Increase difficulty for certain age range
	if girl.age >= 19 && girl.age <= 24 {
		difficulty += 1
	}
	// Increase difficulty based on height and weight thresholds
	if girl.height >= 68 || girl.height <= 63 {
		difficulty += 1
	}
	if girl.weight >= 110 || girl.weight <= 130 {
		difficulty += 1
	}
	// Increase difficulty based on certain hair and eye colors
	if girl.hairColor == Blonde || girl.hairColor == Brown {
		difficulty += 1
	}
	if girl.eyeColor == Blue || girl.eyeColor == Green {
		difficulty += 2
	}
	// Increase difficulty if she is in college or a sorority
	if girl.college {
		difficulty += 1
	}
	if girl.sorority {
		difficulty += 2
	}
	return difficulty
}

// ValidateName checks if a provided name is valid and assigns it to the Girl struct if so.
// If invalid, prompts the user to enter a valid name until a correct one is provided.
//
// Parameters:
//   - name: A pointer to the input string representing the name.
//   - girl: A pointer to the Girl struct where the validated name will be stored.
func ValidateName(name *string, girl *Girl) {
	if isValidName(*name) {
		girl.name = *name
		return
	}
	// Prompt the user to enter a valid name until it meets criteria.
	fmt.Println("Invalid name: must start with uppercase and contain alphabetic characters only")
	for {
		fmt.Println("Please enter a valid name:")
		_, err := fmt.Scanln(name)
		if err != nil {
			continue
		}
		if isValidName(*name) {
			girl.name = *name
			return
		}
	}
}

// isValidName checks if the name matches a specific pattern: capitalized and alphabetic.
//
// Parameters:
//   - name: A string representing the name to validate.
//
// Returns:
//   - bool: Returns true if the name matches the expected pattern, otherwise false.
func isValidName(name string) bool {
	namePattern := `^[A-Z][a-z]*(\s[A-Z][a-z]*)*$`
	return regexp.MustCompile(namePattern).MatchString(strings.TrimSpace(name))
}

// ValidateHairColor checks if the provided hair color is valid (Brown, Blonde, Black, or Red).
// If not, it prompts the user to re-enter a valid hair color until a correct value is provided.
//
// Parameters:
//   - hairColor: A pointer to a string representing the hair color.
//   - girl: A pointer to the Girl struct where the validated hair color will be stored.
func ValidateHairColor(hairColor *string, girl *Girl) {
	if isValidHairColor(*hairColor) {
		girl.hairColor = HairColor(*hairColor)
		return
	}
	for {
		fmt.Println("Please enter a valid hair color: Brown, Blonde, Black, or Red.")
		_, err := fmt.Scanln(hairColor)
		if err != nil {
			continue
		}
		if isValidHairColor(*hairColor) {
			girl.hairColor = HairColor(*hairColor)
			return
		}
	}
}

// isValidHairColor checks if the hair color matches one of the defined constants.
//
// Parameters:
//   - color: A string representing the hair color to validate.
//
// Returns:
//   - bool: Returns true if the color is valid, otherwise false.
func isValidHairColor(color string) bool {
	color = strings.ToLower(color)
	return color == "brown" || color == "blonde" || color == "black" || color == "red"
}

// ValidateEyeColor checks if the provided eye color is valid (Brown, Blue, Black, Green, or Hazel).
// If not, it prompts the user to re-enter a valid eye color until a correct value is provided.
//
// Parameters:
//   - eyeColor: A pointer to a string representing the eye color.
//   - girl: A pointer to the Girl struct where the validated eye color will be stored.
func ValidateEyeColor(eyeColor *string, girl *Girl) {
	if isValidEyeColor(*eyeColor) {
		girl.eyeColor = EyeColor(*eyeColor)
		return
	}
	for {
		fmt.Println("Please enter a valid eye color: Brown, Blue, Black, Green, or Hazel.")
		_, err := fmt.Scanln(eyeColor)
		if err != nil {
			continue
		}
		if isValidEyeColor(*eyeColor) {
			girl.eyeColor = EyeColor(*eyeColor)
			return
		}
	}
}

// isValidEyeColor checks if the eye color matches one of the defined constants.
//
// Parameters:
//   - color: A string representing the eye color to validate.
//
// Returns:
//   - bool: Returns true if the color is valid, otherwise false.
func isValidEyeColor(color string) bool {
	color = strings.ToLower(color)
	return color == "brown" || color == "hazel" || color == "black" || color == "green" || color == "blue"
}
