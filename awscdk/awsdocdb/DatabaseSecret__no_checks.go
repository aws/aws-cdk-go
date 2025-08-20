//go:build no_runtime_type_checking

package awsdocdb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseSecret) validateAddReplicaRegionParameters(region *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateAddRotationScheduleParameters(id *string, options *awssecretsmanager.RotationScheduleOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateAttachParameters(target awssecretsmanager.ISecretAttachmentTarget) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DatabaseSecret) validateSecretValueFromJsonParameters(jsonField *string) error {
	return nil
}

func validateDatabaseSecret_FromSecretAttributesParameters(scope constructs.Construct, id *string, attrs *awssecretsmanager.SecretAttributes) error {
	return nil
}

func validateDatabaseSecret_FromSecretCompleteArnParameters(scope constructs.Construct, id *string, secretCompleteArn *string) error {
	return nil
}

func validateDatabaseSecret_FromSecretNameV2Parameters(scope constructs.Construct, id *string, secretName *string) error {
	return nil
}

func validateDatabaseSecret_FromSecretPartialArnParameters(scope constructs.Construct, id *string, secretPartialArn *string) error {
	return nil
}

func validateDatabaseSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseSecret_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseSecret_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseSecret_IsSecretParameters(x interface{}) error {
	return nil
}

func validateNewDatabaseSecretParameters(scope constructs.Construct, id *string, props *DatabaseSecretProps) error {
	return nil
}

