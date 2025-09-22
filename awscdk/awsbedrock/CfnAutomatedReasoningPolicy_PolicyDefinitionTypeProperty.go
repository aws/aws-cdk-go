package awsbedrock


// A custom type definition within the policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyDefinitionTypeProperty := &PolicyDefinitionTypeProperty{
//   	Name: jsii.String("name"),
//   	Values: []interface{}{
//   		&PolicyDefinitionTypeValueProperty{
//   			Value: jsii.String("value"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html
//
type CfnAutomatedReasoningPolicy_PolicyDefinitionTypeProperty struct {
	// The name of a custom type defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html#cfn-bedrock-automatedreasoningpolicy-policydefinitiontype-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The possible values for a custom type defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html#cfn-bedrock-automatedreasoningpolicy-policydefinitiontype-values
	//
	Values interface{} `field:"required" json:"values" yaml:"values"`
	// A description of the custom type defined in the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html#cfn-bedrock-automatedreasoningpolicy-policydefinitiontype-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

