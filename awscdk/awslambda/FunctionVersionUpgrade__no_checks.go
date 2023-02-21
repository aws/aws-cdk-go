//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionVersionUpgrade) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewFunctionVersionUpgradeParameters(featureFlag *string) error {
	return nil
}

