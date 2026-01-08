//go:build no_runtime_type_checking

package awssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISecretTargetAttachment) validateAddRotationScheduleParameters(id *string, options *RotationScheduleOptions) error {
	return nil
}

func (i *jsiiProxy_ISecretTargetAttachment) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_ISecretTargetAttachment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_ISecretTargetAttachment) validateAttachParameters(target ISecretAttachmentTarget) error {
	return nil
}

func (i *jsiiProxy_ISecretTargetAttachment) validateCfnDynamicReferenceKeyParameters(options *awscdk.SecretsManagerSecretOptions) error {
	return nil
}

func (i *jsiiProxy_ISecretTargetAttachment) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ISecretTargetAttachment) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_ISecretTargetAttachment) validateSecretValueFromJsonParameters(key *string) error {
	return nil
}

