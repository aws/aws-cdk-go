//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Grant) validateCombineParameters(rhs Grant) error {
	return nil
}

func validateGrant_AddToPrincipalParameters(options *GrantOnPrincipalOptions) error {
	return nil
}

func validateGrant_AddToPrincipalAndResourceParameters(options *GrantOnPrincipalAndResourceOptions) error {
	return nil
}

func validateGrant_AddToPrincipalOrResourceParameters(options *GrantWithResourceOptions) error {
	return nil
}

func validateGrant_DropParameters(grantee IGrantable, _intent *string) error {
	return nil
}

