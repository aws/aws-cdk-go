//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsSendMessage) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (s *jsiiProxy_SqsSendMessage) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewSqsSendMessageParameters(queue awssqs.IQueue, props *SqsSendMessageProps) error {
	return nil
}

