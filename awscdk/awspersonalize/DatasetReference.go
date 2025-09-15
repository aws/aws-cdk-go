package awspersonalize


// A reference to a Dataset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetReference := &DatasetReference{
//   	DatasetArn: jsii.String("datasetArn"),
//   }
//
type DatasetReference struct {
	// The DatasetArn of the Dataset resource.
	DatasetArn *string `field:"required" json:"datasetArn" yaml:"datasetArn"`
}

