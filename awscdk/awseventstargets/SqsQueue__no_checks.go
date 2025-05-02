//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsQueue) validateBindParameters(rule awsevents.IRule) error {
	return nil
}

func validateNewSqsQueueParameters(queue awssqs.IQueue, props *SqsQueueProps) error {
	return nil
}

