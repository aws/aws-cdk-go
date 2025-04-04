//go:build no_runtime_type_checking

package awssecretsmanager

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Secret) validateAddReplicaRegionParameters(region *string) error {
	return nil
}

func (s *jsiiProxy_Secret) validateAddRotationScheduleParameters(id *string, options *RotationScheduleOptions) error {
	return nil
}

func (s *jsiiProxy_Secret) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (s *jsiiProxy_Secret) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Secret) validateAttachParameters(target ISecretAttachmentTarget) error {
	return nil
}

func (s *jsiiProxy_Secret) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Secret) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_Secret) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_Secret) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (s *jsiiProxy_Secret) validateSecretValueFromJsonParameters(jsonField *string) error {
	return nil
}

func validateSecret_FromSecretAttributesParameters(scope constructs.Construct, id *string, attrs *SecretAttributes) error {
	return nil
}

func validateSecret_FromSecretCompleteArnParameters(scope constructs.Construct, id *string, secretCompleteArn *string) error {
	return nil
}

func validateSecret_FromSecretNameV2Parameters(scope constructs.Construct, id *string, secretName *string) error {
	return nil
}

func validateSecret_FromSecretPartialArnParameters(scope constructs.Construct, id *string, secretPartialArn *string) error {
	return nil
}

func validateSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSecret_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecret_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSecret_IsSecretParameters(x interface{}) error {
	return nil
}

func validateNewSecretParameters(scope constructs.Construct, id *string, props *SecretProps) error {
	return nil
}

