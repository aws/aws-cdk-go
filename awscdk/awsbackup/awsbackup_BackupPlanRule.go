package awsbackup

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
// Experimental.
type BackupPlanRule interface {
	// Properties of BackupPlanRule.
	// Experimental.
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


// Experimental.
func NewBackupPlanRule(props *BackupPlanRuleProps) BackupPlanRule {
	_init_.Initialize()

	j := jsiiProxy_BackupPlanRule{}

	_jsii_.Create(
		"monocdk.aws_backup.BackupPlanRule",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewBackupPlanRule_Override(b BackupPlanRule, props *BackupPlanRuleProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_backup.BackupPlanRule",
		[]interface{}{props},
		b,
	)
}

// Daily with 35 days retention.
// Experimental.
func BackupPlanRule_Daily(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"daily",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 1 year retention, move to cold storage after 1 month.
// Experimental.
func BackupPlanRule_Monthly1Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"monthly1Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 5 year retention, move to cold storage after 3 months.
// Experimental.
func BackupPlanRule_Monthly5Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"monthly5Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Monthly 7 year retention, move to cold storage after 3 months.
// Experimental.
func BackupPlanRule_Monthly7Year(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"monthly7Year",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

// Weekly with 3 months retention.
// Experimental.
func BackupPlanRule_Weekly(backupVault IBackupVault) BackupPlanRule {
	_init_.Initialize()

	var returns BackupPlanRule

	_jsii_.StaticInvoke(
		"monocdk.aws_backup.BackupPlanRule",
		"weekly",
		[]interface{}{backupVault},
		&returns,
	)

	return returns
}

