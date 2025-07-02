//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IParameterGroup) validateAddParameterParameters(key *string, value *string) error {
	return nil
}

func (i *jsiiProxy_IParameterGroup) validateBindToClusterParameters(options *ParameterGroupClusterBindOptions) error {
	return nil
}

func (i *jsiiProxy_IParameterGroup) validateBindToInstanceParameters(options *ParameterGroupInstanceBindOptions) error {
	return nil
}

