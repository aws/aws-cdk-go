package awsiotsitewise


// The data source for the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   datasetSourceProperty := &DatasetSourceProperty{
//   	SourceFormat: jsii.String("sourceFormat"),
//   	SourceType: jsii.String("sourceType"),
//
//   	// the properties below are optional
//   	SourceDetail: &SourceDetailProperty{
//   		Kendra: &KendraSourceDetailProperty{
//   			KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetsource.html
//
type CfnDataset_DatasetSourceProperty struct {
	// The format of the dataset source associated with the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetsource.html#cfn-iotsitewise-dataset-datasetsource-sourceformat
	//
	SourceFormat *string `field:"required" json:"sourceFormat" yaml:"sourceFormat"`
	// The type of data source for the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetsource.html#cfn-iotsitewise-dataset-datasetsource-sourcetype
	//
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// The details of the dataset source associated with the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetsource.html#cfn-iotsitewise-dataset-datasetsource-sourcedetail
	//
	SourceDetail interface{} `field:"optional" json:"sourceDetail" yaml:"sourceDetail"`
}

