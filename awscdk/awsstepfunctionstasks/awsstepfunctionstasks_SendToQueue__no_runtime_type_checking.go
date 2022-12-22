//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SendToQueue) validateBindParameters(_task awsstepfunctions.Task) error {
	return nil
}

func validateNewSendToQueueParameters(queue awssqs.IQueue, props *SendToQueueProps) error {
	return nil
}

