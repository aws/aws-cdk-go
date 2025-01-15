package awsbackup


// Properties for defining a `CfnLogicallyAirGappedBackupVault`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicy interface{}
//
//   cfnLogicallyAirGappedBackupVaultProps := &CfnLogicallyAirGappedBackupVaultProps{
//   	BackupVaultName: jsii.String("backupVaultName"),
//   	MaxRetentionDays: jsii.Number(123),
//   	MinRetentionDays: jsii.Number(123),
//
//   	// the properties below are optional
//   	AccessPolicy: accessPolicy,
//   	BackupVaultTags: map[string]*string{
//   		"backupVaultTagsKey": jsii.String("backupVaultTags"),
//   	},
//   	Notifications: &NotificationObjectTypeProperty{
//   		BackupVaultEvents: []*string{
//   			jsii.String("backupVaultEvents"),
//   		},
//   		SnsTopicArn: jsii.String("snsTopicArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html
//
type CfnLogicallyAirGappedBackupVaultProps struct {
	// The name of a logical container where backups are stored.
	//
	// Logically air-gapped backup vaults are identified by names that are unique to the account used to create them and the Region where they are created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-backupvaultname
	//
	BackupVaultName *string `field:"required" json:"backupVaultName" yaml:"backupVaultName"`
	// The maximum retention period that the vault retains its recovery points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-maxretentiondays
	//
	MaxRetentionDays *float64 `field:"required" json:"maxRetentionDays" yaml:"maxRetentionDays"`
	// This setting specifies the minimum retention period that the vault retains its recovery points.
	//
	// The minimum value accepted is 7 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-minretentiondays
	//
	MinRetentionDays *float64 `field:"required" json:"minRetentionDays" yaml:"minRetentionDays"`
	// The backup vault access policy document in JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-accesspolicy
	//
	AccessPolicy interface{} `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// The tags to assign to the vault.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-backupvaulttags
	//
	BackupVaultTags *map[string]*string `field:"optional" json:"backupVaultTags" yaml:"backupVaultTags"`
	// Returns event notifications for the specified backup vault.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-notifications
	//
	Notifications interface{} `field:"optional" json:"notifications" yaml:"notifications"`
}

