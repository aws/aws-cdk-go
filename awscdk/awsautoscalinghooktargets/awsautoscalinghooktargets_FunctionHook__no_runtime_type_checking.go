//go:build no_runtime_type_checking

package awsautoscalinghooktargets

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FunctionHook) validateBindParameters(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) error {
	return nil
}

func validateNewFunctionHookParameters(fn awslambda.IFunction) error {
	return nil
}

