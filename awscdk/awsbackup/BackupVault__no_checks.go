//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupVault) validateAddToAccessPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (b *jsiiProxy_BackupVault) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BackupVault) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BackupVault) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (b *jsiiProxy_BackupVault) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateBackupVault_FromBackupVaultArnParameters(scope constructs.Construct, id *string, backupVaultArn *string) error {
	return nil
}

func validateBackupVault_FromBackupVaultNameParameters(scope constructs.Construct, id *string, backupVaultName *string) error {
	return nil
}

func validateBackupVault_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBackupVault_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBackupVault_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBackupVaultParameters(scope constructs.Construct, id *string, props *BackupVaultProps) error {
	return nil
}

