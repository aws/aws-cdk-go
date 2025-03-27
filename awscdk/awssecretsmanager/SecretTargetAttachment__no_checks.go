//go:build no_runtime_type_checking

package awssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretTargetAttachment) validateAddRotationScheduleParameters(id *string, options *RotationScheduleOptions) error {
	return nil
}

func (s *jsiiProxy_SecretTargetAttachment) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_SecretTargetAttachment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SecretTargetAttachment) validateAttachParameters(target ISecretAttachmentTarget) error {
	return nil
}

func (s *jsiiProxy_SecretTargetAttachment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SecretTargetAttachment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_SecretTargetAttachment) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SecretTargetAttachment) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_SecretTargetAttachment) validateSecretValueFromJsonParameters(jsonField *string) error {
	return nil
}

func validateSecretTargetAttachment_FromSecretTargetAttachmentSecretArnParameters(scope constructs.Construct, id *string, secretTargetAttachmentSecretArn *string) error {
	return nil
}

func validateSecretTargetAttachment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecretTargetAttachment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecretTargetAttachment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSecretTargetAttachmentParameters(scope constructs.Construct, id *string, props *SecretTargetAttachmentProps) error {
	return nil
}

