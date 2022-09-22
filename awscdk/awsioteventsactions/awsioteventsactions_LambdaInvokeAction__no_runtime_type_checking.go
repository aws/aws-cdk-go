//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsioteventsactions

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaInvokeAction) validateBindParameters(_scope constructs.Construct, options *awsiotevents.ActionBindOptions) error {
	return nil
}

func validateNewLambdaInvokeActionParameters(func_ awslambda.IFunction) error {
	return nil
}

