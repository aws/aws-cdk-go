package interfacesawspersonalize


// A reference to a DataDeletionJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataDeletionJobReference := &DataDeletionJobReference{
//   	DataDeletionJobArn: jsii.String("dataDeletionJobArn"),
//   }
//
type DataDeletionJobReference struct {
	// The DataDeletionJobArn of the DataDeletionJob resource.
	DataDeletionJobArn *string `field:"required" json:"dataDeletionJobArn" yaml:"dataDeletionJobArn"`
}

