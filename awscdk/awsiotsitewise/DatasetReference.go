package awsiotsitewise


// A reference to a Dataset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetReference := &DatasetReference{
//   	DatasetArn: jsii.String("datasetArn"),
//   	DatasetId: jsii.String("datasetId"),
//   }
//
type DatasetReference struct {
	// The ARN of the Dataset resource.
	DatasetArn *string `field:"required" json:"datasetArn" yaml:"datasetArn"`
	// The DatasetId of the Dataset resource.
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
}

