//go:build no_runtime_type_checking

package awslambdadestinations

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsDestination) validateBindParameters(_scope constructs.Construct, fn awslambda.IFunction, _options *awslambda.DestinationOptions) error {
	return nil
}

func validateNewSnsDestinationParameters(topic awssns.ITopic) error {
	return nil
}

