//go:build no_runtime_type_checking

package awssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISecret) validateAddRotationScheduleParameters(id *string, options *RotationScheduleOptions) error {
	return nil
}

func (i *jsiiProxy_ISecret) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_ISecret) validateAttachParameters(target ISecretAttachmentTarget) error {
	return nil
}

func (i *jsiiProxy_ISecret) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ISecret) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ISecret) validateSecretValueFromJsonParameters(key *string) error {
	return nil
}

