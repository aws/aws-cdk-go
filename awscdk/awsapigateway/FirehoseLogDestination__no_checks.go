//go:build no_runtime_type_checking

package awsapigateway

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehoseLogDestination) validateBindParameters(_stage IStage) error {
	return nil
}

func validateNewFirehoseLogDestinationParameters(stream awskinesisfirehose.CfnDeliveryStream) error {
	return nil
}

