package previewawsbedrockmixins


// A variable defined within the policy that can be used in rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDefinitionVariableProperty := &PolicyDefinitionVariableProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionvariable.html
//
type CfnAutomatedReasoningPolicyPropsMixin_PolicyDefinitionVariableProperty struct {
	// A description of a variable defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionvariable.html#cfn-bedrock-automatedreasoningpolicy-policydefinitionvariable-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of a variable defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionvariable.html#cfn-bedrock-automatedreasoningpolicy-policydefinitionvariable-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The data type of a variable defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionvariable.html#cfn-bedrock-automatedreasoningpolicy-policydefinitionvariable-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

