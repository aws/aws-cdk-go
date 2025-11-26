//go:build no_runtime_type_checking

package previewawslogs

// Building without runtime type checking enabled, so all the below just return nil

func (x *jsiiProxy_XRayLogsDelivery) validateBindParameters(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) error {
	return nil
}

