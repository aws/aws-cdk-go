//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MissingRemovalPolicies) validateApplyParameters(policy RemovalPolicy, props *RemovalPolicyProps) error {
	return nil
}

func (m *jsiiProxy_MissingRemovalPolicies) validateDestroyParameters(props *RemovalPolicyProps) error {
	return nil
}

func (m *jsiiProxy_MissingRemovalPolicies) validateRetainParameters(props *RemovalPolicyProps) error {
	return nil
}

func (m *jsiiProxy_MissingRemovalPolicies) validateRetainOnUpdateOrDeleteParameters(props *RemovalPolicyProps) error {
	return nil
}

func (m *jsiiProxy_MissingRemovalPolicies) validateSnapshotParameters(props *RemovalPolicyProps) error {
	return nil
}

func validateMissingRemovalPolicies_OfParameters(scope constructs.IConstruct) error {
	return nil
}

