//go:build no_runtime_type_checking

package awsautoscalinghooktargets

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TopicHook) validateBindParameters(_scope constructs.Construct, options *awsautoscaling.BindHookTargetOptions) error {
	return nil
}

func validateNewTopicHookParameters(topic awssns.ITopic) error {
	return nil
}

