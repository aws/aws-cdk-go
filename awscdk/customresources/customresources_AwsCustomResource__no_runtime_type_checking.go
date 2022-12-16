//go:build no_runtime_type_checking

package customresources

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsCustomResource) validateGetResponseFieldParameters(dataPath *string) error {
	return nil
}

func (a *jsiiProxy_AwsCustomResource) validateGetResponseFieldReferenceParameters(dataPath *string) error {
	return nil
}

func validateAwsCustomResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAwsCustomResourceParameters(scope constructs.Construct, id *string, props *AwsCustomResourceProps) error {
	return nil
}

