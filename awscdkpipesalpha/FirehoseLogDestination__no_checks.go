//go:build no_runtime_type_checking

package awscdkpipesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehoseLogDestination) validateBindParameters(_pipe IPipe) error {
	return nil
}

func (f *jsiiProxy_FirehoseLogDestination) validateGrantPushParameters(pipeRole awsiam.IRole) error {
	return nil
}

func validateNewFirehoseLogDestinationParameters(deliveryStream awscdkkinesisfirehosealpha.IDeliveryStream) error {
	return nil
}

