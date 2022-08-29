package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A backup plan rule.
//
// Example:
//   var plan backupPlan
//
//   plan.addRule(backup.backupPlanRule.daily())
//   plan.addRule(backup.backupPlanRule.weekly())
//
type BackupPlanRule interface {
	// Properties of BackupPlanRule.
	Props() *BackupPlanRuleProps
}

// The jsii proxy struct for BackupPlanRule
type jsiiProxy_BackupPlanRule struct {
	_ byte // padding
}

func (j *jsiiProxy_BackupPlanRule) Props() *BackupPlanRuleProps {
	var returns *BackupPlanRuleProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


func NewBackupPlanRule(props *BackupPlanRuleProps) BackupPlanRule {
	_init_.Initialize()

	j := jsiiProxy_BackupPlanRule{}

	_jsii_.Create(
		"aws-cdk-lib.aws_backup.BackupPlanRule",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewBackupPlanRule_Override(b BackupPlanRule, props *BackupPlanRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_backup.BackupPlanRule",
		[]interface{}{props},
		b,
	)
}

// Daily with 35 days retention.
func BackupPlanRule_Daily(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupPlanRule",
		"daily",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 1 year retention, move to cold storage after 1 month.
func BackupPlanRule_Monthly1Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupPlanRule",
		"monthly1Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 5 year retention, move to cold storage after 3 months.
func BackupPlanRule_Monthly5Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupPlanRule",
		"monthly5Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 7 year retention, move to cold storage after 3 months.
func BackupPlanRule_Monthly7Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupPlanRule",
		"monthly7Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Weekly with 3 months retention.
func BackupPlanRule_Weekly(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_backup.BackupPlanRule",
		"weekly",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

