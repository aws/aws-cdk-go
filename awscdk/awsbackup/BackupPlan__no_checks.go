//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupPlan) validateAddRuleParameters(rule BackupPlanRule) error {
	return nil
}

func (b *jsiiProxy_BackupPlan) validateAddSelectionParameters(id *string, options *BackupSelectionOptions) error {
	return nil
}

func (b *jsiiProxy_BackupPlan) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (b *jsiiProxy_BackupPlan) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (b *jsiiProxy_BackupPlan) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateBackupPlan_Daily35DayRetentionParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateBackupPlan_DailyMonthly1YearRetentionParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateBackupPlan_DailyWeeklyMonthly5YearRetentionParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateBackupPlan_DailyWeeklyMonthly7YearRetentionParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateBackupPlan_FromBackupPlanIdParameters(scope constructs.Construct, id *string, backupPlanId *string) error {
	return nil
}

func validateBackupPlan_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBackupPlan_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateBackupPlan_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewBackupPlanParameters(scope constructs.Construct, id *string, props *BackupPlanProps) error {
	return nil
}

