//go:build no_runtime_type_checking

package awsautoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ILifecycleHookTarget) validateBindParameters(scope constructs.Construct, options *BindHookTargetOptions) error {
	return nil
}

