//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FirehoseTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (f *jsiiProxy_FirehoseTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewFirehoseTargetParameters(deliveryStream awskinesisfirehose.IDeliveryStream, parameters *FirehoseTargetParameters) error {
	return nil
}

