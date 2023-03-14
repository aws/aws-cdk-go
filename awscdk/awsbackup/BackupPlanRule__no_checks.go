//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func validateNewBackupPlanRuleParameters(props *BackupPlanRuleProps) error {
	return nil
}

