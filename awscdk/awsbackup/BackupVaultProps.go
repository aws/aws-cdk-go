package awsbackup

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Properties for a BackupVault.
//
// Example:
//   myKey := kms.Key_FromKeyArn(this, jsii.String("MyKey"), jsii.String("aaa"))
//   myTopic := sns.Topic_FromTopicArn(this, jsii.String("MyTopic"), jsii.String("bbb"))
//
//   vault := backup.NewBackupVault(this, jsii.String("Vault"), &BackupVaultProps{
//   	EncryptionKey: myKey,
//   	 // Custom encryption key
//   	NotificationTopic: myTopic,
//   })
//
type BackupVaultProps struct {
	// A resource-based policy that is used to manage access permissions on the backup vault.
	AccessPolicy awsiam.PolicyDocument `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// The name of a logical container where backups are stored.
	//
	// Backup vaults
	// are identified by names that are unique to the account used to create
	// them and the AWS Region where they are created.
	BackupVaultName *string `field:"optional" json:"backupVaultName" yaml:"backupVaultName"`
	// Whether to add statements to the vault access policy that prevents anyone from deleting a recovery point.
	BlockRecoveryPointDeletion *bool `field:"optional" json:"blockRecoveryPointDeletion" yaml:"blockRecoveryPointDeletion"`
	// The server-side encryption key to use to protect your backups.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Configuration for AWS Backup Vault Lock.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/vault-lock.html
	//
	LockConfiguration *LockConfiguration `field:"optional" json:"lockConfiguration" yaml:"lockConfiguration"`
	// The vault events to send.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/sns-notifications.html
	//
	NotificationEvents *[]BackupVaultEvents `field:"optional" json:"notificationEvents" yaml:"notificationEvents"`
	// A SNS topic to send vault events to.
	// See: https://docs.aws.amazon.com/aws-backup/latest/devguide/sns-notifications.html
	//
	NotificationTopic awssns.ITopic `field:"optional" json:"notificationTopic" yaml:"notificationTopic"`
	// The removal policy to apply to the vault.
	//
	// Note that removing a vault
	// that contains recovery points will fail.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
}

