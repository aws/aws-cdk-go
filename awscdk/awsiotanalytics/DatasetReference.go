package awsiotanalytics


// A reference to a Dataset resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetReference := &DatasetReference{
//   	DatasetName: jsii.String("datasetName"),
//   }
//
type DatasetReference struct {
	// The DatasetName of the Dataset resource.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
}

