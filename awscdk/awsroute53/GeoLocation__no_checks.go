//go:build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func validateGeoLocation_ContinentParameters(continentCode Continent) error {
	return nil
}

func validateGeoLocation_CountryParameters(countryCode *string) error {
	return nil
}

func validateGeoLocation_SubdivisionParameters(subdivisionCode *string) error {
	return nil
}

