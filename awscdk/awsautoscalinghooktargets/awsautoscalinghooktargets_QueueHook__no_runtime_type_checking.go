//go:build no_runtime_type_checking

package awsautoscalinghooktargets

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueHook) validateBindParameters(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) error {
	return nil
}

func validateNewQueueHookParameters(queue awssqs.IQueue) error {
	return nil
}

