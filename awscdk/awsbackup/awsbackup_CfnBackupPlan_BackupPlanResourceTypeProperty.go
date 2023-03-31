package awsbackup


// Specifies an object containing properties used to create a backup plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var backupOptions interface{}
//
//   backupPlanResourceTypeProperty := &backupPlanResourceTypeProperty{
//   	backupPlanName: jsii.String("backupPlanName"),
//   	backupPlanRule: []interface{}{
//   		&backupRuleResourceTypeProperty{
//   			ruleName: jsii.String("ruleName"),
//   			targetBackupVault: jsii.String("targetBackupVault"),
//
//   			// the properties below are optional
//   			completionWindowMinutes: jsii.Number(123),
//   			copyActions: []interface{}{
//   				&copyActionResourceTypeProperty{
//   					destinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//
//   					// the properties below are optional
//   					lifecycle: &lifecycleResourceTypeProperty{
//   						deleteAfterDays: jsii.Number(123),
//   						moveToColdStorageAfterDays: jsii.Number(123),
//   					},
//   				},
//   			},
//   			enableContinuousBackup: jsii.Boolean(false),
//   			lifecycle: &lifecycleResourceTypeProperty{
//   				deleteAfterDays: jsii.Number(123),
//   				moveToColdStorageAfterDays: jsii.Number(123),
//   			},
//   			recoveryPointTags: map[string]*string{
//   				"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   			},
//   			scheduleExpression: jsii.String("scheduleExpression"),
//   			startWindowMinutes: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	advancedBackupSettings: []interface{}{
//   		&advancedBackupSettingResourceTypeProperty{
//   			backupOptions: backupOptions,
//   			resourceType: jsii.String("resourceType"),
//   		},
//   	},
//   }
//
type CfnBackupPlan_BackupPlanResourceTypeProperty struct {
	// The display name of a backup plan.
	BackupPlanName *string `field:"required" json:"backupPlanName" yaml:"backupPlanName"`
	// An array of `BackupRule` objects, each of which specifies a scheduled task that is used to back up a selection of resources.
	BackupPlanRule interface{} `field:"required" json:"backupPlanRule" yaml:"backupPlanRule"`
	// A list of backup options for each resource type.
	AdvancedBackupSettings interface{} `field:"optional" json:"advancedBackupSettings" yaml:"advancedBackupSettings"`
}

