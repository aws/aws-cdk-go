//go:build no_runtime_type_checking

package awscdkpipessourcesalpha

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisSource) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (k *jsiiProxy_KinesisSource) validateGetDeadLetterTargetArnParameters(deadLetterTarget interface{}) error {
	return nil
}

func (k *jsiiProxy_KinesisSource) validateGrantPushParameters(grantee awsiam.IRole, deadLetterTarget interface{}) error {
	return nil
}

func (k *jsiiProxy_KinesisSource) validateGrantReadParameters(grantee awsiam.IRole) error {
	return nil
}

func validateKinesisSource_IsSourceWithDeadLetterTargetParameters(source awscdkpipesalpha.ISource) error {
	return nil
}

func validateNewKinesisSourceParameters(stream awskinesis.IStream, parameters *KinesisSourceParameters) error {
	return nil
}

