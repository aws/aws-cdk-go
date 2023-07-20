package awsbackup


// Properties for defining a `CfnBackupVault`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var accessPolicy interface{}
//
//   cfnBackupVaultProps := &CfnBackupVaultProps{
//   	BackupVaultName: jsii.String("backupVaultName"),
//
//   	// the properties below are optional
//   	AccessPolicy: accessPolicy,
//   	BackupVaultTags: map[string]*string{
//   		"backupVaultTagsKey": jsii.String("backupVaultTags"),
//   	},
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	LockConfiguration: &LockConfigurationTypeProperty{
//   		MinRetentionDays: jsii.Number(123),
//
//   		// the properties below are optional
//   		ChangeableForDays: jsii.Number(123),
//   		MaxRetentionDays: jsii.Number(123),
//   	},
//   	Notifications: &NotificationObjectTypeProperty{
//   		BackupVaultEvents: []*string{
//   			jsii.String("backupVaultEvents"),
//   		},
//   		SnsTopicArn: jsii.String("snsTopicArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html
//
type CfnBackupVaultProps struct {
	// The name of a logical container where backups are stored.
	//
	// Backup vaults are identified by names that are unique to the account used to create them and the AWS Region where they are created. They consist of lowercase letters, numbers, and hyphens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaultname
	//
	BackupVaultName *string `field:"required" json:"backupVaultName" yaml:"backupVaultName"`
	// A resource-based policy that is used to manage access permissions on the target backup vault.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-accesspolicy
	//
	AccessPolicy interface{} `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// Metadata that you can assign to help organize the resources that you create.
	//
	// Each tag is a key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-backupvaulttags
	//
	BackupVaultTags interface{} `field:"optional" json:"backupVaultTags" yaml:"backupVaultTags"`
	// A server-side encryption key you can specify to encrypt your backups from services that support full AWS Backup management;
	//
	// for example, `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` . If you specify a key, you must specify its ARN, not its alias. If you do not specify a key, AWS Backup creates a KMS key for you by default.
	//
	// To learn which AWS Backup services support full AWS Backup management and how AWS Backup handles encryption for backups from services that do not yet support full AWS Backup , see [Encryption for backups in AWS Backup](https://docs.aws.amazon.com/aws-backup/latest/devguide/encryption.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// Configuration for [AWS Backup Vault Lock](https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-lockconfiguration
	//
	LockConfiguration interface{} `field:"optional" json:"lockConfiguration" yaml:"lockConfiguration"`
	// The SNS event notifications for the specified backup vault.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-backup-backupvault.html#cfn-backup-backupvault-notifications
	//
	Notifications interface{} `field:"optional" json:"notifications" yaml:"notifications"`
}

