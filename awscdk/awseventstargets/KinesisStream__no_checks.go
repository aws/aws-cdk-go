//go:build no_runtime_type_checking

package awseventstargets

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisStream) validateBindParameters(_rule awsevents.IRule) error {
	return nil
}

func validateNewKinesisStreamParameters(stream awskinesis.IStream, props *KinesisStreamProps) error {
	return nil
}

