package mixinsawsiotanalytics


// Information about the dataset whose content generation triggers the new dataset content generation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   triggeringDatasetProperty := &TriggeringDatasetProperty{
//   	DatasetName: jsii.String("datasetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-triggeringdataset.html
//
type CfnDatasetPropsMixin_TriggeringDatasetProperty struct {
	// The name of the data set whose content generation triggers the new data set content generation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-triggeringdataset.html#cfn-iotanalytics-dataset-triggeringdataset-datasetname
	//
	DatasetName *string `field:"optional" json:"datasetName" yaml:"datasetName"`
}

