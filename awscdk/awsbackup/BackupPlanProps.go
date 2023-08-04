package awsbackup


// Properties for a BackupPlan.
//
// Example:
//   plan := backup.NewBackupPlan(this, jsii.String("Plan"), &BackupPlanProps{
//   	WindowsVss: jsii.Boolean(true),
//   })
//
type BackupPlanProps struct {
	// The display name of the backup plan.
	// Default: - A CDK generated name.
	//
	BackupPlanName *string `field:"optional" json:"backupPlanName" yaml:"backupPlanName"`
	// Rules for the backup plan.
	//
	// Use `addRule()` to add rules after
	// instantiation.
	// Default: - use `addRule()` to add rules.
	//
	BackupPlanRules *[]BackupPlanRule `field:"optional" json:"backupPlanRules" yaml:"backupPlanRules"`
	// The backup vault where backups are stored.
	// Default: - use the vault defined at the rule level. If not defined a new
	// common vault for the plan will be created.
	//
	BackupVault IBackupVault `field:"optional" json:"backupVault" yaml:"backupVault"`
	// Enable Windows VSS backup.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/windows-backups.html
	//
	// Default: false.
	//
	WindowsVss *bool `field:"optional" json:"windowsVss" yaml:"windowsVss"`
}

