//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBridgeTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (e *jsiiProxy_EventBridgeTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewEventBridgeTargetParameters(eventBus awsevents.IEventBus, parameters *EventBridgeTargetParameters) error {
	return nil
}

