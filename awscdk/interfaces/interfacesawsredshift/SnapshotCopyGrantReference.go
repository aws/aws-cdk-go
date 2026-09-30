package interfacesawsredshift


// A reference to a SnapshotCopyGrant resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   snapshotCopyGrantReference := &SnapshotCopyGrantReference{
//   	SnapshotCopyGrantArn: jsii.String("snapshotCopyGrantArn"),
//   }
//
type SnapshotCopyGrantReference struct {
	// The Arn of the SnapshotCopyGrant resource.
	SnapshotCopyGrantArn *string `field:"required" json:"snapshotCopyGrantArn" yaml:"snapshotCopyGrantArn"`
}

