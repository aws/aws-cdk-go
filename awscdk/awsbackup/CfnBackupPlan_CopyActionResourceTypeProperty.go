package awsbackup


// Copies backups created by a backup rule to another vault.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   copyActionResourceTypeProperty := &CopyActionResourceTypeProperty{
//   	DestinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//
//   	// the properties below are optional
//   	Lifecycle: &LifecycleResourceTypeProperty{
//   		DeleteAfterDays: jsii.Number(123),
//   		MoveToColdStorageAfterDays: jsii.Number(123),
//   		OptInToArchiveForSupportedResources: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-copyactionresourcetype.html
//
type CfnBackupPlan_CopyActionResourceTypeProperty struct {
	// An Amazon Resource Name (ARN) that uniquely identifies the destination backup vault for the copied backup.
	//
	// For example, `arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-copyactionresourcetype.html#cfn-backup-backupplan-copyactionresourcetype-destinationbackupvaultarn
	//
	DestinationBackupVaultArn *string `field:"required" json:"destinationBackupVaultArn" yaml:"destinationBackupVaultArn"`
	// Defines when a protected resource is transitioned to cold storage and when it expires.
	//
	// AWS Backup transitions and expires backups automatically according to the lifecycle that you define. If you do not specify a lifecycle, AWS Backup applies the lifecycle policy of the source backup to the destination backup.
	//
	// Backups transitioned to cold storage must be stored in cold storage for a minimum of 90 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-backup-backupplan-copyactionresourcetype.html#cfn-backup-backupplan-copyactionresourcetype-lifecycle
	//
	Lifecycle interface{} `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

