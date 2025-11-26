//go:build no_runtime_type_checking

package previewawslogs

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehoseLogsDelivery) validateBindParameters(scope constructs.IConstruct, deliverySource interfacesawslogs.IDeliverySourceRef, sourceResourceArn *string) error {
	return nil
}

func validateNewFirehoseLogsDeliveryParameters(stream interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

