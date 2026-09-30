package interfacesawsfsx


// A reference to a Backup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupReference := &BackupReference{
//   	BackupId: jsii.String("backupId"),
//   	ResourceArn: jsii.String("resourceArn"),
//   }
//
type BackupReference struct {
	// The BackupId of the Backup resource.
	BackupId *string `field:"required" json:"backupId" yaml:"backupId"`
	// The ResourceARN of the Backup resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

