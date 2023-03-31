package awsiotanalytics


// Information about the dataset whose content generation triggers the new dataset content generation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   triggeringDatasetProperty := &triggeringDatasetProperty{
//   	datasetName: jsii.String("datasetName"),
//   }
//
type CfnDataset_TriggeringDatasetProperty struct {
	// The name of the data set whose content generation triggers the new data set content generation.
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
}

