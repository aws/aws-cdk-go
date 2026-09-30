package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicIdentifierDeclarationProperty := &TopicIdentifierDeclarationProperty{
//   	Identifier: jsii.String("identifier"),
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topicidentifierdeclaration.html
//
type CfnAnalysis_TopicIdentifierDeclarationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topicidentifierdeclaration.html#cfn-quicksight-analysis-topicidentifierdeclaration-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-topicidentifierdeclaration.html#cfn-quicksight-analysis-topicidentifierdeclaration-topicarn
	//
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

