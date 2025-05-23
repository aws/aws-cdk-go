//go:build no_runtime_type_checking

package awsappsync

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBridgeDataSource) validateCreateFunctionParameters(id *string, props *BaseAppsyncFunctionProps) error {
	return nil
}

func (e *jsiiProxy_EventBridgeDataSource) validateCreateResolverParameters(id *string, props *BaseResolverProps) error {
	return nil
}

func validateEventBridgeDataSource_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_EventBridgeDataSource) validateSetApiParameters(val IGraphqlApi) error {
	return nil
}

func validateNewEventBridgeDataSourceParameters(scope constructs.Construct, id *string, props *EventBridgeDataSourceProps) error {
	return nil
}

