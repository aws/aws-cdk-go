//go:build no_runtime_type_checking

package awslogsdestinations

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisDestination) validateBindParameters(scope constructs.Construct, _sourceLogGroup awslogs.ILogGroup) error {
	return nil
}

func validateNewKinesisDestinationParameters(stream awskinesis.IStream, props *KinesisDestinationProps) error {
	return nil
}

