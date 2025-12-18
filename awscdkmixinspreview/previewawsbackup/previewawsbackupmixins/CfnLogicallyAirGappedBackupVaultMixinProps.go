package previewawsbackupmixins


// Properties for CfnLogicallyAirGappedBackupVaultPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var accessPolicy interface{}
//
//   cfnLogicallyAirGappedBackupVaultMixinProps := &CfnLogicallyAirGappedBackupVaultMixinProps{
//   	AccessPolicy: accessPolicy,
//   	BackupVaultName: jsii.String("backupVaultName"),
//   	BackupVaultTags: map[string]*string{
//   		"backupVaultTagsKey": jsii.String("backupVaultTags"),
//   	},
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	MaxRetentionDays: jsii.Number(123),
//   	MinRetentionDays: jsii.Number(123),
//   	MpaApprovalTeamArn: jsii.String("mpaApprovalTeamArn"),
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
type CfnLogicallyAirGappedBackupVaultMixinProps struct {
	// The backup vault access policy document in JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-accesspolicy
	//
	AccessPolicy interface{} `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// The name of a logical container where backups are stored.
	//
	// Logically air-gapped backup vaults are identified by names that are unique to the account used to create them and the Region where they are created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-backupvaultname
	//
	BackupVaultName *string `field:"optional" json:"backupVaultName" yaml:"backupVaultName"`
	// The tags to assign to the vault.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-backupvaulttags
	//
	BackupVaultTags *map[string]*string `field:"optional" json:"backupVaultTags" yaml:"backupVaultTags"`
	// The server-side encryption key that is used to protect your backups; for example, `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` .
	//
	// If this field is left blank, AWS Backup will create an AWS owned key to be used to encrypt the content of the logically air-gapped vault. The ARN of this created key will be available as `Fn::GetAtt` output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The maximum retention period that the vault retains its recovery points.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-maxretentiondays
	//
	MaxRetentionDays *float64 `field:"optional" json:"maxRetentionDays" yaml:"maxRetentionDays"`
	// This setting specifies the minimum retention period that the vault retains its recovery points.
	//
	// The minimum value accepted is 7 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-minretentiondays
	//
	MinRetentionDays *float64 `field:"optional" json:"minRetentionDays" yaml:"minRetentionDays"`
	// The Amazon Resource Name (ARN) of the MPA approval team to associate with the backup vault.
	//
	// This cannot be changed after it is set from the CloudFormation template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-mpaapprovalteamarn
	//
	MpaApprovalTeamArn *string `field:"optional" json:"mpaApprovalTeamArn" yaml:"mpaApprovalTeamArn"`
	// Returns event notifications for the specified backup vault.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-logicallyairgappedbackupvault.html#cfn-backup-logicallyairgappedbackupvault-notifications
	//
	Notifications interface{} `field:"optional" json:"notifications" yaml:"notifications"`
}

