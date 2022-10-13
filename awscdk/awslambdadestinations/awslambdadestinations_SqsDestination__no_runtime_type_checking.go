//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awslambdadestinations

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsDestination) validateBindParameters(_scope awscdk.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) error {
	return nil
}

func validateNewSqsDestinationParameters(queue awssqs.IQueue) error {
	return nil
}

