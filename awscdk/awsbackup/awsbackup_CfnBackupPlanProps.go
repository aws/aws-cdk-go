package awsbackup


// Properties for defining a `CfnBackupPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var backupOptions interface{}
//
//   cfnBackupPlanProps := &cfnBackupPlanProps{
//   	backupPlan: &backupPlanResourceTypeProperty{
//   		backupPlanName: jsii.String("backupPlanName"),
//   		backupPlanRule: []interface{}{
//   			&backupRuleResourceTypeProperty{
//   				ruleName: jsii.String("ruleName"),
//   				targetBackupVault: jsii.String("targetBackupVault"),
//
//   				// the properties below are optional
//   				completionWindowMinutes: jsii.Number(123),
//   				copyActions: []interface{}{
//   					&copyActionResourceTypeProperty{
//   						destinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//
//   						// the properties below are optional
//   						lifecycle: &lifecycleResourceTypeProperty{
//   							deleteAfterDays: jsii.Number(123),
//   							moveToColdStorageAfterDays: jsii.Number(123),
//   						},
//   					},
//   				},
//   				enableContinuousBackup: jsii.Boolean(false),
//   				lifecycle: &lifecycleResourceTypeProperty{
//   					deleteAfterDays: jsii.Number(123),
//   					moveToColdStorageAfterDays: jsii.Number(123),
//   				},
//   				recoveryPointTags: map[string]*string{
//   					"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   				},
//   				scheduleExpression: jsii.String("scheduleExpression"),
//   				startWindowMinutes: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		advancedBackupSettings: []interface{}{
//   			&advancedBackupSettingResourceTypeProperty{
//   				backupOptions: backupOptions,
//   				resourceType: jsii.String("resourceType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	backupPlanTags: map[string]*string{
//   		"backupPlanTagsKey": jsii.String("backupPlanTags"),
//   	},
//   }
//
type CfnBackupPlanProps struct {
	// Uniquely identifies the backup plan to be associated with the selection of resources.
	BackupPlan interface{} `field:"required" json:"backupPlan" yaml:"backupPlan"`
	// To help organize your resources, you can assign your own metadata to the resources that you create.
	//
	// Each tag is a key-value pair. The specified tags are assigned to all backups created with this plan.
	BackupPlanTags interface{} `field:"optional" json:"backupPlanTags" yaml:"backupPlanTags"`
}

