//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IPrincipal) validateAddToPrincipalPolicyParameters(statement PolicyStatement) error {
	return nil
}

