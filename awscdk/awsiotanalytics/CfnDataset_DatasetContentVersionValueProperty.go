package awsiotanalytics


// The dataset whose latest contents are used as input to the notebook or application.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetContentVersionValueProperty := &DatasetContentVersionValueProperty{
//   	DatasetName: jsii.String("datasetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-datasetcontentversionvalue.html
//
type CfnDataset_DatasetContentVersionValueProperty struct {
	// The name of the dataset whose latest contents are used as input to the notebook or application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-datasetcontentversionvalue.html#cfn-iotanalytics-dataset-datasetcontentversionvalue-datasetname
	//
	DatasetName *string `field:"required" json:"datasetName" yaml:"datasetName"`
}

