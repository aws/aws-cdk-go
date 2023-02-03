package awsbackup


// Properties for a BackupPlan.
//
// Example:
//   plan := backup.NewBackupPlan(this, jsii.String("Plan"), &backupPlanProps{
//   	windowsVss: jsii.Boolean(true),
//   })
//
type BackupPlanProps struct {
	// The display name of the backup plan.
	BackupPlanName *string `field:"optional" json:"backupPlanName" yaml:"backupPlanName"`
	// Rules for the backup plan.
	//
	// Use `addRule()` to add rules after
	// instantiation.
	BackupPlanRules *[]BackupPlanRule `field:"optional" json:"backupPlanRules" yaml:"backupPlanRules"`
	// The backup vault where backups are stored.
	BackupVault IBackupVault `field:"optional" json:"backupVault" yaml:"backupVault"`
	// Enable Windows VSS backup.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/windows-backups.html
	//
	WindowsVss *bool `field:"optional" json:"windowsVss" yaml:"windowsVss"`
}

