//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UnknownPrincipal) validateAddToPolicyParameters(statement PolicyStatement) error {
	return nil
}

func (u *jsiiProxy_UnknownPrincipal) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

func validateNewUnknownPrincipalParameters(props *UnknownPrincipalProps) error {
	return nil
}

