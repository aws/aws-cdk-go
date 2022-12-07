//go:build no_runtime_type_checking

package integtests

// Building without runtime type checking enabled, so all the below just return nil

func validateActualResult_FromAwsApiCallParameters(query IAwsApiCall, attribute *string) error {
	return nil
}

func validateActualResult_FromCustomResourceParameters(customResource awscdk.CustomResource, attribute *string) error {
	return nil
}

func (j *jsiiProxy_ActualResult) validateSetResultParameters(val *string) error {
	return nil
}

