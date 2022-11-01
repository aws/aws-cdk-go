//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func validateNewApiEventSourceParameters(method *string, path *string, options *awsapigateway.MethodOptions) error {
	return nil
}

