package previewawswisdommixins


// An object that can be used to specify tag conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   tagConditionProperty := &TagConditionProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tagcondition.html
//
type CfnAIAgentPropsMixin_TagConditionProperty struct {
	// The tag key in the tag condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tagcondition.html#cfn-wisdom-aiagent-tagcondition-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value in the tag condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tagcondition.html#cfn-wisdom-aiagent-tagcondition-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

