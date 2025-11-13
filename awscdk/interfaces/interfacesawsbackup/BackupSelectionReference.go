package interfacesawsbackup


// A reference to a BackupSelection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupSelectionReference := &BackupSelectionReference{
//   	BackupSelectionId: jsii.String("backupSelectionId"),
//   }
//
type BackupSelectionReference struct {
	// The Id of the BackupSelection resource.
	BackupSelectionId *string `field:"required" json:"backupSelectionId" yaml:"backupSelectionId"`
}

