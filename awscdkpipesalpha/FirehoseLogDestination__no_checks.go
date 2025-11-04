//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehoseLogDestination) validateBindParameters(pipe IPipe) error {
	return nil
}

func (f *jsiiProxy_FirehoseLogDestination) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewFirehoseLogDestinationParameters(deliveryStream awskinesisfirehose.IDeliveryStream) error {
	return nil
}

