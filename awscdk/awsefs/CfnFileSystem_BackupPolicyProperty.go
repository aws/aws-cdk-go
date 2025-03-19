package awsefs


// The backup policy turns automatic backups for the file system on or off.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupPolicyProperty := &BackupPolicyProperty{
//   	Status: jsii.String("status"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-backuppolicy.html
//
type CfnFileSystem_BackupPolicyProperty struct {
	// Set the backup policy status for the file system.
	//
	// - *`ENABLED`* - Turns automatic backups on for the file system.
	// - *`DISABLED`* - Turns automatic backups off for the file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-filesystem-backuppolicy.html#cfn-efs-filesystem-backuppolicy-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
}

