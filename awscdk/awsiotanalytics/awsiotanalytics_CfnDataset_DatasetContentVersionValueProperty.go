package awsiotanalytics


// The dataset whose latest contents are used as input to the notebook or application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentVersionValueProperty := &datasetContentVersionValueProperty{
//   	datasetName: jsii.String("datasetName"),
//   }
//
type CfnDataset_DatasetContentVersionValueProperty struct {
	// The name of the dataset whose latest contents are used as input to the notebook or application.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
}

