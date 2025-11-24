//go:build no_runtime_type_checking

package core

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConstructSelector) validateSelectParameters(scope constructs.IConstruct) error {
	return nil
}

func validateConstructSelector_ByIdParameters(pattern interface{}) error {
	return nil
}

func validateConstructSelector_ResourcesOfTypeParameters(type_ interface{}) error {
	return nil
}

