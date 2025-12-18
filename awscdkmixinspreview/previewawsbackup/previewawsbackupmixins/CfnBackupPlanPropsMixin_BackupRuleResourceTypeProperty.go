package previewawsbackupmixins


// Specifies an object containing properties used to schedule a task to back up a selection of resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   backupRuleResourceTypeProperty := &BackupRuleResourceTypeProperty{
//   	CompletionWindowMinutes: jsii.Number(123),
//   	CopyActions: []interface{}{
//   		&CopyActionResourceTypeProperty{
//   			DestinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//   			Lifecycle: &LifecycleResourceTypeProperty{
//   				DeleteAfterDays: jsii.Number(123),
//   				MoveToColdStorageAfterDays: jsii.Number(123),
//   				OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	EnableContinuousBackup: jsii.Boolean(false),
//   	IndexActions: []interface{}{
//   		&IndexActionsResourceTypeProperty{
//   			ResourceTypes: []*string{
//   				jsii.String("resourceTypes"),
//   			},
//   		},
//   	},
//   	Lifecycle: &LifecycleResourceTypeProperty{
//   		DeleteAfterDays: jsii.Number(123),
//   		MoveToColdStorageAfterDays: jsii.Number(123),
//   		OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   	},
//   	RecoveryPointTags: map[string]*string{
//   		"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   	},
//   	RuleName: jsii.String("ruleName"),
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	ScheduleExpressionTimezone: jsii.String("scheduleExpressionTimezone"),
//   	StartWindowMinutes: jsii.Number(123),
//   	TargetBackupVault: jsii.String("targetBackupVault"),
//   	TargetLogicallyAirGappedBackupVaultArn: jsii.String("targetLogicallyAirGappedBackupVaultArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html
//
type CfnBackupPlanPropsMixin_BackupRuleResourceTypeProperty struct {
	// A value in minutes after a backup job is successfully started before it must be completed or it is canceled by AWS Backup .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-completionwindowminutes
	//
	CompletionWindowMinutes *float64 `field:"optional" json:"completionWindowMinutes" yaml:"completionWindowMinutes"`
	// An array of CopyAction objects, which contains the details of the copy operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-copyactions
	//
	CopyActions interface{} `field:"optional" json:"copyActions" yaml:"copyActions"`
	// Enables continuous backup and point-in-time restores (PITR).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-enablecontinuousbackup
	//
	EnableContinuousBackup interface{} `field:"optional" json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// There can up to one IndexAction in each BackupRule, as each backup can have 0 or 1 backup index associated with it.
	//
	// Within the array is ResourceTypes. Only 1 resource type will be accepted for each BackupRule. Valid values:
	//
	// - `EBS` for Amazon Elastic Block Store
	// - `S3` for Amazon Simple Storage Service (Amazon S3).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-indexactions
	//
	IndexActions interface{} `field:"optional" json:"indexActions" yaml:"indexActions"`
	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.
	//
	// AWS Backup transitions and expires backups automatically according to the lifecycle that you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-lifecycle
	//
	Lifecycle interface{} `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// The tags to assign to the resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-recoverypointtags
	//
	RecoveryPointTags interface{} `field:"optional" json:"recoveryPointTags" yaml:"recoveryPointTags"`
	// A display name for a backup rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-scheduleexpression
	//
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// This is the timezone in which the schedule expression is set.
	//
	// By default, ScheduleExpressions are in UTC. You can modify this to a specified timezone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-scheduleexpressiontimezone
	//
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// An optional value that specifies a period of time in minutes after a backup is scheduled before a job is canceled if it doesn't start successfully.
	//
	// If this value is included, it must be at least 60 minutes to avoid errors.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-startwindowminutes
	//
	StartWindowMinutes *float64 `field:"optional" json:"startWindowMinutes" yaml:"startWindowMinutes"`
	// The name of a logical container where backups are stored.
	//
	// Backup vaults are identified by names that are unique to the account used to create them and the AWS Region where they are created. They consist of letters, numbers, and hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-targetbackupvault
	//
	TargetBackupVault *string `field:"optional" json:"targetBackupVault" yaml:"targetBackupVault"`
	// The ARN of a logically air-gapped vault.
	//
	// ARN must be in the same account and Region. If provided, supported fully managed resources back up directly to logically air-gapped vault, while other supported resources create a temporary (billable) snapshot in backup vault, then copy it to logically air-gapped vault. Unsupported resources only back up to the specified backup vault.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-backupruleresourcetype.html#cfn-backup-backupplan-backupruleresourcetype-targetlogicallyairgappedbackupvaultarn
	//
	TargetLogicallyAirGappedBackupVaultArn *string `field:"optional" json:"targetLogicallyAirGappedBackupVaultArn" yaml:"targetLogicallyAirGappedBackupVaultArn"`
}

