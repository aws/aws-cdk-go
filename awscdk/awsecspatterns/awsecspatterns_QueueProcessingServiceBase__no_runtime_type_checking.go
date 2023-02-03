//go:build no_runtime_type_checking

package awsecspatterns

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueProcessingServiceBase) validateConfigureAutoscalingForServiceParameters(service awsecs.BaseService) error {
	return nil
}

func (q *jsiiProxy_QueueProcessingServiceBase) validateGetDefaultClusterParameters(scope constructs.Construct) error {
	return nil
}

func (q *jsiiProxy_QueueProcessingServiceBase) validateGrantPermissionsToServiceParameters(service awsecs.BaseService) error {
	return nil
}

func validateQueueProcessingServiceBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewQueueProcessingServiceBaseParameters(scope constructs.Construct, id *string, props *QueueProcessingServiceBaseProps) error {
	return nil
}

