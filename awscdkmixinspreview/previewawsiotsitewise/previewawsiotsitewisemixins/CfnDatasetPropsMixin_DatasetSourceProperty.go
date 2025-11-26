package previewawsiotsitewisemixins


// The data source for the dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   datasetSourceProperty := &DatasetSourceProperty{
//   	SourceDetail: &SourceDetailProperty{
//   		Kendra: &KendraSourceDetailProperty{
//   			KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   			RoleArn: jsii.String("roleArn"),
//   		},
//   	},
//   	SourceFormat: jsii.String("sourceFormat"),
//   	SourceType: jsii.String("sourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetsource.html
//
type CfnDatasetPropsMixin_DatasetSourceProperty struct {
	// The details of the dataset source associated with the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetsource.html#cfn-iotsitewise-dataset-datasetsource-sourcedetail
	//
	SourceDetail interface{} `field:"optional" json:"sourceDetail" yaml:"sourceDetail"`
	// The format of the dataset source associated with the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetsource.html#cfn-iotsitewise-dataset-datasetsource-sourceformat
	//
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
	// The type of data source for the dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-datasetsource.html#cfn-iotsitewise-dataset-datasetsource-sourcetype
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

