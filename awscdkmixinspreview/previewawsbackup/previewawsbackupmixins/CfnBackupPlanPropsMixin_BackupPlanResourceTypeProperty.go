package previewawsbackupmixins


// Specifies an object containing properties used to create a backup plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var backupOptions interface{}
//
//   backupPlanResourceTypeProperty := &BackupPlanResourceTypeProperty{
//   	AdvancedBackupSettings: []interface{}{
//   		&AdvancedBackupSettingResourceTypeProperty{
//   			BackupOptions: backupOptions,
//   			ResourceType: jsii.String("resourceType"),
//   		},
//   	},
//   	BackupPlanName: jsii.String("backupPlanName"),
//   	BackupPlanRule: []interface{}{
//   		&BackupRuleResourceTypeProperty{
//   			CompletionWindowMinutes: jsii.Number(123),
//   			CopyActions: []interface{}{
//   				&CopyActionResourceTypeProperty{
//   					DestinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//   					Lifecycle: &LifecycleResourceTypeProperty{
//   						DeleteAfterDays: jsii.Number(123),
//   						MoveToColdStorageAfterDays: jsii.Number(123),
//   						OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			EnableContinuousBackup: jsii.Boolean(false),
//   			IndexActions: []interface{}{
//   				&IndexActionsResourceTypeProperty{
//   					ResourceTypes: []*string{
//   						jsii.String("resourceTypes"),
//   					},
//   				},
//   			},
//   			Lifecycle: &LifecycleResourceTypeProperty{
//   				DeleteAfterDays: jsii.Number(123),
//   				MoveToColdStorageAfterDays: jsii.Number(123),
//   				OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   			},
//   			RecoveryPointTags: map[string]*string{
//   				"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   			},
//   			RuleName: jsii.String("ruleName"),
//   			ScanActions: []interface{}{
//   				&ScanActionResourceTypeProperty{
//   					MalwareScanner: jsii.String("malwareScanner"),
//   					ScanMode: jsii.String("scanMode"),
//   				},
//   			},
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//   			ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   			StartWindowMinutes: jsii.Number(123),
//   			TargetBackupVault: jsii.String("targetBackupVault"),
//   			TargetLogicallyAirGappedBackupVaultArn: jsii.String("targetLogicallyAirGappedBackupVaultArn"),
//   		},
//   	},
//   	ScanSettings: []interface{}{
//   		&ScanSettingResourceTypeProperty{
//   			MalwareScanner: jsii.String("malwareScanner"),
//   			ResourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   			ScannerRoleArn: jsii.String("scannerRoleArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupplanresourcetype.html
//
type CfnBackupPlanPropsMixin_BackupPlanResourceTypeProperty struct {
	// A list of backup options for each resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupplanresourcetype.html#cfn-backup-backupplan-backupplanresourcetype-advancedbackupsettings
	//
	AdvancedBackupSettings interface{} `field:"optional" json:"advancedBackupSettings" yaml:"advancedBackupSettings"`
	// The display name of a backup plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupplanresourcetype.html#cfn-backup-backupplan-backupplanresourcetype-backupplanname
	//
	BackupPlanName *string `field:"optional" json:"backupPlanName" yaml:"backupPlanName"`
	// An array of `BackupRule` objects, each of which specifies a scheduled task that is used to back up a selection of resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupplanresourcetype.html#cfn-backup-backupplan-backupplanresourcetype-backupplanrule
	//
	BackupPlanRule interface{} `field:"optional" json:"backupPlanRule" yaml:"backupPlanRule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupplanresourcetype.html#cfn-backup-backupplan-backupplanresourcetype-scansettings
	//
	ScanSettings interface{} `field:"optional" json:"scanSettings" yaml:"scanSettings"`
}

