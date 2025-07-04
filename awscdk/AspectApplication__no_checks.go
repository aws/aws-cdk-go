//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_AspectApplication) validateSetPriorityParameters(val *float64) error {
	return nil
}

func validateNewAspectApplicationParameters(construct constructs.IConstruct, aspect IAspect, priority *float64) error {
	return nil
}

