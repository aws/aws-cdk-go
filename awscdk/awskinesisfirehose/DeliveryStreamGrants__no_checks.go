//go:build no_runtime_type_checking

package awskinesisfirehose

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DeliveryStreamGrants) validatePutRecordsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateDeliveryStreamGrants_FromDeliveryStreamParameters(resource interfacesawskinesisfirehose.IDeliveryStreamRef) error {
	return nil
}

