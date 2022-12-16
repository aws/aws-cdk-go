//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduledTaskBase) validateAddTaskAsTargetParameters(ecsTaskTarget awseventstargets.EcsTask) error {
	return nil
}

func (s *jsiiProxy_ScheduledTaskBase) validateAddTaskDefinitionToEventTargetParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (s *jsiiProxy_ScheduledTaskBase) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (s *jsiiProxy_ScheduledTaskBase) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateScheduledTaskBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewScheduledTaskBaseParameters(scope constructs.Construct, id *string, props *ScheduledTaskBaseProps) error {
	return nil
}

