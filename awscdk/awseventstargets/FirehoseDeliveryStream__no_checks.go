//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehoseDeliveryStream) validateBindParameters(rule interfacesawsevents.IRuleRef) error {
	return nil
}

func validateNewFirehoseDeliveryStreamParameters(deliveryStream awskinesisfirehose.IDeliveryStream, props *FirehoseDeliveryStreamProps) error {
	return nil
}

