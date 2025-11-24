package mixinsawsquicksight


// Configuration options for a `Topic` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   topicConfigOptionsProperty := &TopicConfigOptionsProperty{
//   	QBusinessInsightsEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicconfigoptions.html
//
type CfnTopicPropsMixin_TopicConfigOptionsProperty struct {
	// Enables Amazon Q Business Insights for a `Topic` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-topicconfigoptions.html#cfn-quicksight-topic-topicconfigoptions-qbusinessinsightsenabled
	//
	QBusinessInsightsEnabled interface{} `field:"optional" json:"qBusinessInsightsEnabled" yaml:"qBusinessInsightsEnabled"`
}

