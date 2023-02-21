//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueProcessingFargateService) validateConfigureAutoscalingForServiceParameters(service awsecs.BaseService) error {
	return nil
}

func (q *jsiiProxy_QueueProcessingFargateService) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (q *jsiiProxy_QueueProcessingFargateService) validateGrantPermissionsToServiceParameters(service awsecs.BaseService) error {
	return nil
}

func validateQueueProcessingFargateService_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewQueueProcessingFargateServiceParameters(scope constructs.Construct, id *string, props *QueueProcessingFargateServiceProps) error {
	return nil
}

