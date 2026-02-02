package previewawsbackupmixins


// Properties for CfnBackupPlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var backupOptions interface{}
//
//   cfnBackupPlanMixinProps := &CfnBackupPlanMixinProps{
//   	BackupPlan: &BackupPlanResourceTypeProperty{
//   		AdvancedBackupSettings: []interface{}{
//   			&AdvancedBackupSettingResourceTypeProperty{
//   				BackupOptions: backupOptions,
//   				ResourceType: jsii.String("resourceType"),
//   			},
//   		},
//   		BackupPlanName: jsii.String("backupPlanName"),
//   		BackupPlanRule: []interface{}{
//   			&BackupRuleResourceTypeProperty{
//   				CompletionWindowMinutes: jsii.Number(123),
//   				CopyActions: []interface{}{
//   					&CopyActionResourceTypeProperty{
//   						DestinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//   						Lifecycle: &LifecycleResourceTypeProperty{
//   							DeleteAfterDays: jsii.Number(123),
//   							MoveToColdStorageAfterDays: jsii.Number(123),
//   							OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   						},
//   					},
//   				},
//   				EnableContinuousBackup: jsii.Boolean(false),
//   				IndexActions: []interface{}{
//   					&IndexActionsResourceTypeProperty{
//   						ResourceTypes: []*string{
//   							jsii.String("resourceTypes"),
//   						},
//   					},
//   				},
//   				Lifecycle: &LifecycleResourceTypeProperty{
//   					DeleteAfterDays: jsii.Number(123),
//   					MoveToColdStorageAfterDays: jsii.Number(123),
//   					OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   				},
//   				RecoveryPointTags: map[string]*string{
//   					"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   				},
//   				RuleName: jsii.String("ruleName"),
//   				ScanActions: []interface{}{
//   					&ScanActionResourceTypeProperty{
//   						MalwareScanner: jsii.String("malwareScanner"),
//   						ScanMode: jsii.String("scanMode"),
//   					},
//   				},
//   				ScheduleExpression: jsii.String("scheduleExpression"),
//   				ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   				StartWindowMinutes: jsii.Number(123),
//   				TargetBackupVault: jsii.String("targetBackupVault"),
//   				TargetLogicallyAirGappedBackupVaultArn: jsii.String("targetLogicallyAirGappedBackupVaultArn"),
//   			},
//   		},
//   		ScanSettings: []interface{}{
//   			&ScanSettingResourceTypeProperty{
//   				MalwareScanner: jsii.String("malwareScanner"),
//   				ResourceTypes: []*string{
//   					jsii.String("resourceTypes"),
//   				},
//   				ScannerRoleArn: jsii.String("scannerRoleArn"),
//   			},
//   		},
//   	},
//   	BackupPlanTags: map[string]*string{
//   		"backupPlanTagsKey": jsii.String("backupPlanTags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html
//
type CfnBackupPlanMixinProps struct {
	// Uniquely identifies the backup plan to be associated with the selection of resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html#cfn-backup-backupplan-backupplan
	//
	BackupPlan interface{} `field:"optional" json:"backupPlan" yaml:"backupPlan"`
	// The tags to assign to the backup plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupplan.html#cfn-backup-backupplan-backupplantags
	//
	BackupPlanTags *map[string]*string `field:"optional" json:"backupPlanTags" yaml:"backupPlanTags"`
}

