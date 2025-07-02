//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IRuleTarget) validateBindParameters(rule IRule) error {
	return nil
}

