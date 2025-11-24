package mixinsawsbedrock


// A value associated with a custom type in the policy definition.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDefinitionTypeValueProperty := &PolicyDefinitionTypeValueProperty{
//   	Description: jsii.String("description"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontypevalue.html
//
type CfnAutomatedReasoningPolicyPropsMixin_PolicyDefinitionTypeValueProperty struct {
	// A description of the policy definition type value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontypevalue.html#cfn-bedrock-automatedreasoningpolicy-policydefinitiontypevalue-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The value associated with a policy definition type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontypevalue.html#cfn-bedrock-automatedreasoningpolicy-policydefinitiontypevalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

