//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ICfnResourceOptions) validateSetCreationPolicyParameters(val *CfnCreationPolicy) error {
	return nil
}

func (j *jsiiProxy_ICfnResourceOptions) validateSetUpdatePolicyParameters(val *CfnUpdatePolicy) error {
	return nil
}

