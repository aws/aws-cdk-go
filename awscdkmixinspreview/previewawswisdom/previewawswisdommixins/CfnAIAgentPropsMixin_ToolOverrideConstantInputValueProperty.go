package previewawswisdommixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   toolOverrideConstantInputValueProperty := &ToolOverrideConstantInputValueProperty{
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideconstantinputvalue.html
//
type CfnAIAgentPropsMixin_ToolOverrideConstantInputValueProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideconstantinputvalue.html#cfn-wisdom-aiagent-tooloverrideconstantinputvalue-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-tooloverrideconstantinputvalue.html#cfn-wisdom-aiagent-tooloverrideconstantinputvalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

