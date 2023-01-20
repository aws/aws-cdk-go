//go:build no_runtime_type_checking

package awslambdadestinations

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBridgeDestination) validateBindParameters(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) error {
	return nil
}

