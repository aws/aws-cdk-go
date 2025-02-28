//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupSelection) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BackupSelection) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BackupSelection) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBackupSelection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBackupSelection_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBackupSelection_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBackupSelectionParameters(scope constructs.Construct, id *string, props *BackupSelectionProps) error {
	return nil
}

