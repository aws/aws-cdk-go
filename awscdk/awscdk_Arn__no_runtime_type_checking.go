//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateArn_ExtractResourceNameParameters(arn *string, resourceType *string) error {
	return nil
}

func validateArn_FormatParameters(components *ArnComponents) error {
	return nil
}

func validateArn_ParseParameters(arn *string) error {
	return nil
}

func validateArn_SplitParameters(arn *string, arnFormat ArnFormat) error {
	return nil
}

