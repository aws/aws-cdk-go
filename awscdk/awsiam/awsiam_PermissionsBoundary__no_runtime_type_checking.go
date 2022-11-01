//go:build no_runtime_type_checking

package awsiam

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PermissionsBoundary) validateApplyParameters(boundaryPolicy IManagedPolicy) error {
	return nil
}

func validatePermissionsBoundary_OfParameters(scope constructs.IConstruct) error {
	return nil
}

