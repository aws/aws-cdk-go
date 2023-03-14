//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateToken_AsAnyParameters(value interface{}) error {
	return nil
}

func validateToken_AsListParameters(value interface{}, options *EncodingOptions) error {
	return nil
}

func validateToken_AsNumberParameters(value interface{}) error {
	return nil
}

func validateToken_AsStringParameters(value interface{}, options *EncodingOptions) error {
	return nil
}

func validateToken_CompareStringsParameters(possibleToken1 *string, possibleToken2 *string) error {
	return nil
}

func validateToken_IsUnresolvedParameters(obj interface{}) error {
	return nil
}

