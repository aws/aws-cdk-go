//go:build no_runtime_type_checking

package awslambdaeventsources

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsEventSource) validateBindParameters(target awslambda.IFunction) error {
	return nil
}

func validateNewSqsEventSourceParameters(queue awssqs.IQueue, props *SqsEventSourceProps) error {
	return nil
}

