package awsiotsitewise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceDetailProperty := &SourceDetailProperty{
//   	Kendra: &KendraSourceDetailProperty{
//   		KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-sourcedetail.html
//
type CfnDataset_SourceDetailProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-dataset-sourcedetail.html#cfn-iotsitewise-dataset-sourcedetail-kendra
	//
	Kendra interface{} `field:"optional" json:"kendra" yaml:"kendra"`
}

