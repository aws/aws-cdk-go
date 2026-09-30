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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topicidentifierdeclaration.html
//
type CfnDashboard_TopicIdentifierDeclarationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topicidentifierdeclaration.html#cfn-quicksight-dashboard-topicidentifierdeclaration-identifier
	//
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-topicidentifierdeclaration.html#cfn-quicksight-dashboard-topicidentifierdeclaration-topicarn
	//
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

