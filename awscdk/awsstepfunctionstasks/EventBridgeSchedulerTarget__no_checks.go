//go:build no_runtime_type_checking

package awsstepfunctionstasks

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_EventBridgeSchedulerTarget) validateSetArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EventBridgeSchedulerTarget) validateSetRetryPolicyParameters(val *RetryPolicy) error {
	return nil
}

func (j *jsiiProxy_EventBridgeSchedulerTarget) validateSetRoleParameters(val awsiam.IRole) error {
	return nil
}

func validateNewEventBridgeSchedulerTargetParameters(props *EventBridgeSchedulerTargetProps) error {
	return nil
}

