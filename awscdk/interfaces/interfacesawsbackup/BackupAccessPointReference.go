package interfacesawsbackup


// A reference to a BackupAccessPoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupAccessPointReference := &BackupAccessPointReference{
//   	BackupAccessPointArn: jsii.String("backupAccessPointArn"),
//   }
//
type BackupAccessPointReference struct {
	// The Arn of the BackupAccessPoint resource.
	BackupAccessPointArn *string `field:"required" json:"backupAccessPointArn" yaml:"backupAccessPointArn"`
}

