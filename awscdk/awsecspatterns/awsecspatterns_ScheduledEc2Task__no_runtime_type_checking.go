//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduledEc2Task) validateAddTaskAsTargetParameters(ecsTaskTarget awseventstargets.EcsTask) error {
	return nil
}

func (s *jsiiProxy_ScheduledEc2Task) validateAddTaskDefinitionToEventTargetParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (s *jsiiProxy_ScheduledEc2Task) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (s *jsiiProxy_ScheduledEc2Task) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateScheduledEc2Task_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewScheduledEc2TaskParameters(scope constructs.Construct, id *string, props *ScheduledEc2TaskProps) error {
	return nil
}

