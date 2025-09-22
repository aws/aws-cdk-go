package awsbackup


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   notificationObjectTypeProperty := &NotificationObjectTypeProperty{
//   	BackupVaultEvents: []*string{
//   		jsii.String("backupVaultEvents"),
//   	},
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-logicallyairgappedbackupvault-notificationobjecttype.html
//
type CfnLogicallyAirGappedBackupVault_NotificationObjectTypeProperty struct {
	// An array of events that indicate the status of jobs to back up resources to the backup vault.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-logicallyairgappedbackupvault-notificationobjecttype.html#cfn-backup-logicallyairgappedbackupvault-notificationobjecttype-backupvaultevents
	//
	BackupVaultEvents *[]*string `field:"required" json:"backupVaultEvents" yaml:"backupVaultEvents"`
	// The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s events;
	//
	// for example, `arn:aws:sns:us-west-2:111122223333:MyVaultTopic` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-logicallyairgappedbackupvault-notificationobjecttype.html#cfn-backup-logicallyairgappedbackupvault-notificationobjecttype-snstopicarn
	//
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

