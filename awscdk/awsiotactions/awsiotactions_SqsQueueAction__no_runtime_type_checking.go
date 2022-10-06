//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsQueueAction) validateBindParameters(rule awsiot.ITopicRule) error {
	return nil
}

func validateNewSqsQueueActionParameters(queue awssqs.IQueue, props *SqsQueueActionProps) error {
	return nil
}

