//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ICfnResourceOptions) validateSetCreationPolicyParameters(val *CfnCreationPolicy) error {
	return nil
}

func (j *jsiiProxy_ICfnResourceOptions) validateSetUpdatePolicyParameters(val *CfnUpdatePolicy) error {
	return nil
}

