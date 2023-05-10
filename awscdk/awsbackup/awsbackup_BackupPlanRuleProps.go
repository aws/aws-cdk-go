package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
)

// Properties for a BackupPlanRule.
//
// Example:
//   var plan backupPlan
//
//   plan.AddRule(backup.NewBackupPlanRule(&BackupPlanRuleProps{
//   	CompletionWindow: awscdk.Duration_Hours(jsii.Number(2)),
//   	StartWindow: awscdk.Duration_*Hours(jsii.Number(1)),
//   	ScheduleExpression: events.Schedule_Cron(&CronOptions{
//   		 // Only cron expressions are supported
//   		Day: jsii.String("15"),
//   		Hour: jsii.String("3"),
//   		Minute: jsii.String("30"),
//   	}),
//   	MoveToColdStorageAfter: awscdk.Duration_Days(jsii.Number(30)),
//   }))
//
// Experimental.
type BackupPlanRuleProps struct {
	// The backup vault where backups are.
	// Experimental.
	BackupVault IBackupVault `field:"optional" json:"backupVault" yaml:"backupVault"`
	// The duration after a backup job is successfully started before it must be completed or it is canceled by AWS Backup.
	// Experimental.
	CompletionWindow awscdk.Duration `field:"optional" json:"completionWindow" yaml:"completionWindow"`
	// Specifies the duration after creation that a recovery point is deleted.
	//
	// Must be greater than `moveToColdStorageAfter`.
	// Experimental.
	DeleteAfter awscdk.Duration `field:"optional" json:"deleteAfter" yaml:"deleteAfter"`
	// Enables continuous backup and point-in-time restores (PITR).
	//
	// Property `deleteAfter` defines the retention period for the backup. It is mandatory if PITR is enabled.
	// If no value is specified, the retention period is set to 35 days which is the maximum retention period supported by PITR.
	//
	// Property `moveToColdStorageAfter` must not be specified because PITR does not support this option.
	// Experimental.
	EnableContinuousBackup *bool `field:"optional" json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// Specifies the duration after creation that a recovery point is moved to cold storage.
	// Experimental.
	MoveToColdStorageAfter awscdk.Duration `field:"optional" json:"moveToColdStorageAfter" yaml:"moveToColdStorageAfter"`
	// A display name for the backup rule.
	// Experimental.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	// Experimental.
	ScheduleExpression awsevents.Schedule `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// The duration after a backup is scheduled before a job is canceled if it doesn't start successfully.
	// Experimental.
	StartWindow awscdk.Duration `field:"optional" json:"startWindow" yaml:"startWindow"`
}

