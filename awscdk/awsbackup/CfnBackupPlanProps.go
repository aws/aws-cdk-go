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
//   cfnBackupPlanProps := &CfnBackupPlanProps{
//   	BackupPlan: &BackupPlanResourceTypeProperty{
//   		BackupPlanName: jsii.String("backupPlanName"),
//   		BackupPlanRule: []interface{}{
//   			&BackupRuleResourceTypeProperty{
//   				RuleName: jsii.String("ruleName"),
//   				TargetBackupVault: jsii.String("targetBackupVault"),
//
//   				// the properties below are optional
//   				CompletionWindowMinutes: jsii.Number(123),
//   				CopyActions: []interface{}{
//   					&CopyActionResourceTypeProperty{
//   						DestinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//
//   						// the properties below are optional
//   						Lifecycle: &LifecycleResourceTypeProperty{
//   							DeleteAfterDays: jsii.Number(123),
//   							MoveToColdStorageAfterDays: jsii.Number(123),
//   							OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				EnableContinuousBackup: jsii.Boolean(false),
//   				Lifecycle: &LifecycleResourceTypeProperty{
//   					DeleteAfterDays: jsii.Number(123),
//   					MoveToColdStorageAfterDays: jsii.Number(123),
//   					OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   				},
//   				RecoveryPointTags: map[string]*string{
//   					"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   				},
//   				ScheduleExpression: jsii.String("scheduleExpression"),
//   				ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   				StartWindowMinutes: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		AdvancedBackupSettings: []interface{}{
//   			&AdvancedBackupSettingResourceTypeProperty{
//   				BackupOptions: backupOptions,
//   				ResourceType: jsii.String("resourceType"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	BackupPlanTags: map[string]*string{
//   		"backupPlanTagsKey": jsii.String("backupPlanTags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html
//
type CfnBackupPlanProps struct {
	// Uniquely identifies the backup plan to be associated with the selection of resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html#cfn-backup-backupplan-backupplan
	//
	BackupPlan interface{} `field:"required" json:"backupPlan" yaml:"backupPlan"`
	// To help organize your resources, you can assign your own metadata to the resources that you create.
	//
	// Each tag is a key-value pair. The specified tags are assigned to all backups created with this plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html#cfn-backup-backupplan-backupplantags
	//
	BackupPlanTags *map[string]*string `field:"optional" json:"backupPlanTags" yaml:"backupPlanTags"`
}

