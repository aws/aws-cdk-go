package interfacesawsrds


// A reference to a DBSnapshot resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBSnapshotReference := &DBSnapshotReference{
//   	DbSnapshotArn: jsii.String("dbSnapshotArn"),
//   }
//
type DBSnapshotReference struct {
	// The DBSnapshotArn of the DBSnapshot resource.
	DbSnapshotArn *string `field:"required" json:"dbSnapshotArn" yaml:"dbSnapshotArn"`
}

