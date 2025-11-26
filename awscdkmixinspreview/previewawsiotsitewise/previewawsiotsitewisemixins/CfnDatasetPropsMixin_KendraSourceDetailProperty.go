package previewawsiotsitewisemixins


// The source details for the Kendra dataset source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   kendraSourceDetailProperty := &KendraSourceDetailProperty{
//   	KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-kendrasourcedetail.html
//
type CfnDatasetPropsMixin_KendraSourceDetailProperty struct {
	// The `knowledgeBaseArn` details for the Kendra dataset source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-kendrasourcedetail.html#cfn-iotsitewise-dataset-kendrasourcedetail-knowledgebasearn
	//
	KnowledgeBaseArn *string `field:"optional" json:"knowledgeBaseArn" yaml:"knowledgeBaseArn"`
	// The `roleARN` details for the Kendra dataset source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-kendrasourcedetail.html#cfn-iotsitewise-dataset-kendrasourcedetail-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

