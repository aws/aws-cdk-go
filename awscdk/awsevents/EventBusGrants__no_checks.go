//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBusGrants) validateAllPutEventsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateEventBusGrants_FromEventBusParameters(resource interfacesawsevents.IEventBusRef) error {
	return nil
}

