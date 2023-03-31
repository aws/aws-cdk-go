package awsbackup


// Specifies an object containing SNS event notification properties for the target backup vault.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationObjectTypeProperty := &notificationObjectTypeProperty{
//   	backupVaultEvents: []*string{
//   		jsii.String("backupVaultEvents"),
//   	},
//   	snsTopicArn: jsii.String("snsTopicArn"),
//   }
//
type CfnBackupVault_NotificationObjectTypeProperty struct {
	// An array of events that indicate the status of jobs to back up resources to the backup vault.
	//
	// For valid events, see [BackupVaultEvents](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_PutBackupVaultNotifications.html#API_PutBackupVaultNotifications_RequestSyntax) in the *AWS Backup API Guide* .
	BackupVaultEvents *[]*string `field:"required" json:"backupVaultEvents" yaml:"backupVaultEvents"`
	// An ARN that uniquely identifies an Amazon Simple Notification Service (Amazon SNS) topic;
	//
	// for example, `arn:aws:sns:us-west-2:111122223333:MyTopic` .
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

