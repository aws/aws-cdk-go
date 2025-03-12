//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceAccount) validateAddToPrincipalPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func validateServiceAccount_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewServiceAccountParameters(scope constructs.Construct, id *string, props *ServiceAccountProps) error {
	return nil
}

