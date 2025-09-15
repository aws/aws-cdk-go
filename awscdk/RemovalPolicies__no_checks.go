//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RemovalPolicies) validateApplyParameters(policy RemovalPolicy, props *RemovalPolicyProps) error {
	return nil
}

func (r *jsiiProxy_RemovalPolicies) validateDestroyParameters(props *RemovalPolicyProps) error {
	return nil
}

func (r *jsiiProxy_RemovalPolicies) validateRetainParameters(props *RemovalPolicyProps) error {
	return nil
}

func (r *jsiiProxy_RemovalPolicies) validateRetainOnUpdateOrDeleteParameters(props *RemovalPolicyProps) error {
	return nil
}

func (r *jsiiProxy_RemovalPolicies) validateSnapshotParameters(props *RemovalPolicyProps) error {
	return nil
}

func validateRemovalPolicies_OfParameters(scope constructs.IConstruct) error {
	return nil
}

