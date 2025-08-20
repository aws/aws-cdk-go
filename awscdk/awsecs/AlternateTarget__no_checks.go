//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlternateTarget) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

func validateNewAlternateTargetParameters(id *string, props *AlternateTargetProps) error {
	return nil
}

