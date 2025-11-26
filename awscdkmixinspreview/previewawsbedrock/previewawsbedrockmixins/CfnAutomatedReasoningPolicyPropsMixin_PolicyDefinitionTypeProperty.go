package previewawsbedrockmixins


// A custom type definition within the policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDefinitionTypeProperty := &PolicyDefinitionTypeProperty{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Values: []interface{}{
//   		&PolicyDefinitionTypeValueProperty{
//   			Description: jsii.String("description"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html
//
type CfnAutomatedReasoningPolicyPropsMixin_PolicyDefinitionTypeProperty struct {
	// A description of the custom type defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html#cfn-bedrock-automatedreasoningpolicy-policydefinitiontype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of a custom type defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html#cfn-bedrock-automatedreasoningpolicy-policydefinitiontype-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The possible values for a custom type defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html#cfn-bedrock-automatedreasoningpolicy-policydefinitiontype-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}

