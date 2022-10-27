//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IEventBus) validateArchiveParameters(id *string, props *BaseArchiveProps) error {
	return nil
}

func (i *jsiiProxy_IEventBus) validateGrantPutEventsToParameters(grantee awsiam.IGrantable) error {
	return nil
}

