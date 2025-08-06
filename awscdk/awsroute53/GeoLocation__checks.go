//go:build !no_runtime_type_checking

package awsroute53

import (
	"fmt"
)

func validateGeoLocation_ContinentParameters(continentCode Continent) error {
	if continentCode == "" {
		return fmt.Errorf("parameter continentCode is required, but nil was provided")
	}

	return nil
}

func validateGeoLocation_CountryParameters(countryCode *string) error {
	if countryCode == nil {
		return fmt.Errorf("parameter countryCode is required, but nil was provided")
	}

	return nil
}

func validateGeoLocation_SubdivisionParameters(subdivisionCode *string) error {
	if subdivisionCode == nil {
		return fmt.Errorf("parameter subdivisionCode is required, but nil was provided")
	}

	return nil
}

