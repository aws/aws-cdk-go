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
//   	VaultState: jsii.String("vaultState"),
//   	VaultType: jsii.String("vaultType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html
//
type CfnLogicallyAirGappedBackupVaultProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-backupvaultname
	//
	BackupVaultName *string `field:"required" json:"backupVaultName" yaml:"backupVaultName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-maxretentiondays
	//
	MaxRetentionDays *float64 `field:"required" json:"maxRetentionDays" yaml:"maxRetentionDays"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-minretentiondays
	//
	MinRetentionDays *float64 `field:"required" json:"minRetentionDays" yaml:"minRetentionDays"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-accesspolicy
	//
	AccessPolicy interface{} `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-backupvaulttags
	//
	BackupVaultTags *map[string]*string `field:"optional" json:"backupVaultTags" yaml:"backupVaultTags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-notifications
	//
	Notifications interface{} `field:"optional" json:"notifications" yaml:"notifications"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-vaultstate
	//
	VaultState *string `field:"optional" json:"vaultState" yaml:"vaultState"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-vaulttype
	//
	VaultType *string `field:"optional" json:"vaultType" yaml:"vaultType"`
}

