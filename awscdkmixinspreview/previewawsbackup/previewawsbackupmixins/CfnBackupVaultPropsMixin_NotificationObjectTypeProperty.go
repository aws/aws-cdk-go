package previewawsbackupmixins


// Specifies an object containing SNS event notification properties for the target backup vault.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   notificationObjectTypeProperty := &NotificationObjectTypeProperty{
//   	BackupVaultEvents: []*string{
//   		jsii.String("backupVaultEvents"),
//   	},
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupvault-notificationobjecttype.html
//
type CfnBackupVaultPropsMixin_NotificationObjectTypeProperty struct {
	// An array of events that indicate the status of jobs to back up resources to the backup vault.
	//
	// For valid events, see [BackupVaultEvents](https://docs.aws.amazon.com/aws-backup/latest/devguide/API_PutBackupVaultNotifications.html#API_PutBackupVaultNotifications_RequestSyntax) in the *AWS Backup API Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupvault-notificationobjecttype.html#cfn-backup-backupvault-notificationobjecttype-backupvaultevents
	//
	BackupVaultEvents *[]*string `field:"optional" json:"backupVaultEvents" yaml:"backupVaultEvents"`
	// An ARN that uniquely identifies an Amazon Simple Notification Service (Amazon SNS) topic;
	//
	// for example, `arn:aws:sns:us-west-2:111122223333:MyTopic` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupvault-notificationobjecttype.html#cfn-backup-backupvault-notificationobjecttype-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

