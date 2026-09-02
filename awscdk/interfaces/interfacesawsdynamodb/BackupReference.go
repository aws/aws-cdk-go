package interfacesawsdynamodb


// A reference to a Backup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupReference := &BackupReference{
//   	BackupArn: jsii.String("backupArn"),
//   }
//
type BackupReference struct {
	// The BackupArn of the Backup resource.
	BackupArn *string `field:"required" json:"backupArn" yaml:"backupArn"`
}

