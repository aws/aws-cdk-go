//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (k *jsiiProxy_KinesisTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewKinesisTargetParameters(stream awskinesis.IStream, parameters *KinesisTargetParameters) error {
	return nil
}

