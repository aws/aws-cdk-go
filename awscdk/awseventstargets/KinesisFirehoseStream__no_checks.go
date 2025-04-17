//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisFirehoseStream) validateBindParameters(_rule awsevents.IRule) error {
	return nil
}

func validateNewKinesisFirehoseStreamParameters(stream awskinesisfirehose.CfnDeliveryStream, props *KinesisFirehoseStreamProps) error {
	return nil
}

