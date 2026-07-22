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
//   	DataDeletionJobId: jsii.String("dataDeletionJobId"),
//   }
//
type DataDeletionJobReference struct {
	// The ARN of the DataDeletionJob resource.
	DataDeletionJobArn *string `field:"required" json:"dataDeletionJobArn" yaml:"dataDeletionJobArn"`
	// The Id of the DataDeletionJob resource.
	DataDeletionJobId *string `field:"required" json:"dataDeletionJobId" yaml:"dataDeletionJobId"`
}

