//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisFirehoseStreamV2) validateBindParameters(rule interfacesawsevents.IRuleRef) error {
	return nil
}

func validateNewKinesisFirehoseStreamV2Parameters(stream IDeliveryStream, props *KinesisFirehoseStreamProps) error {
	return nil
}

