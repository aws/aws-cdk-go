//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduledFargateTask) validateAddTaskAsTargetParameters(ecsTaskTarget awseventstargets.EcsTask) error {
	return nil
}

func (s *jsiiProxy_ScheduledFargateTask) validateAddTaskDefinitionToEventTargetParameters(taskDefinition awsecs.TaskDefinition) error {
	return nil
}

func (s *jsiiProxy_ScheduledFargateTask) validateCreateAWSLogDriverParameters(prefix *string) error {
	return nil
}

func (s *jsiiProxy_ScheduledFargateTask) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func validateScheduledFargateTask_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewScheduledFargateTaskParameters(scope constructs.Construct, id *string, props *ScheduledFargateTaskProps) error {
	return nil
}

