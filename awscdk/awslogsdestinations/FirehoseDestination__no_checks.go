//go:build no_runtime_type_checking

package awslogsdestinations

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehoseDestination) validateBindParameters(scope constructs.Construct, sourceLogGroup awslogs.ILogGroup) error {
	return nil
}

func validateNewFirehoseDestinationParameters(stream awskinesisfirehose.IDeliveryStream, props *FirehoseDestinationProps) error {
	return nil
}

