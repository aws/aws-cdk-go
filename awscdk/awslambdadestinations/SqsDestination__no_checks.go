//go:build no_runtime_type_checking

package awslambdadestinations

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsDestination) validateBindParameters(scope constructs.Construct, fn awslambda.IFunction, options *awslambda.DestinationOptions) error {
	return nil
}

func validateNewSqsDestinationParameters(queue awssqs.IQueue) error {
	return nil
}

