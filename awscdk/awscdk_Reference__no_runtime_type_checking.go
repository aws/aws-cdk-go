//go:build no_runtime_type_checking

// An experiment to bundle the entire CDK into a single module
package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Reference) validateNewErrorParameters(message *string) error {
	return nil
}

func (r *jsiiProxy_Reference) validateResolveParameters(_context IResolveContext) error {
	return nil
}

func validateReference_IsReferenceParameters(x interface{}) error {
	return nil
}

func validateNewReferenceParameters(value interface{}, target IConstruct) error {
	return nil
}

