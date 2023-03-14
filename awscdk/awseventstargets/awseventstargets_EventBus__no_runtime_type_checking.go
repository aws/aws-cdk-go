//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBus) validateBindParameters(rule awsevents.IRule) error {
	return nil
}

func validateNewEventBusParameters(eventBus awsevents.IEventBus, props *EventBusProps) error {
	return nil
}

