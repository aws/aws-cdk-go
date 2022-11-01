//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func validateNewSnsEventSourceParameters(topic awssns.ITopic, props *SnsEventSourceProps) error {
	return nil
}

