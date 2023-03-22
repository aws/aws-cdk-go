package awsefs


// The backup policy turns automatic backups for the file system on or off.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupPolicyProperty := &backupPolicyProperty{
//   	status: jsii.String("status"),
//   }
//
type CfnFileSystem_BackupPolicyProperty struct {
	// Set the backup policy status for the file system.
	//
	// - *`ENABLED`* - Turns automatic backups on for the file system.
	// - *`DISABLED`* - Turns automatic backups off for the file system.
	Status *string `field:"required" json:"status" yaml:"status"`
}

