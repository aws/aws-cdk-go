//go:build no_runtime_type_checking

// Version 2 of the AWS Cloud Development Kit library
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateArn_ExtractResourceNameParameters(arn *string, resourceType *string) error {
	return nil
}

func validateArn_FormatParameters(components *ArnComponents) error {
	return nil
}

func validateArn_SplitParameters(arn *string, arnFormat ArnFormat) error {
	return nil
}

