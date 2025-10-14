package awsbedrock


// The complete policy definition containing rules, variables, and types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyDefinitionProperty := &PolicyDefinitionProperty{
//   	Rules: []interface{}{
//   		&PolicyDefinitionRuleProperty{
//   			Expression: jsii.String("expression"),
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			AlternateExpression: jsii.String("alternateExpression"),
//   		},
//   	},
//   	Types: []interface{}{
//   		&PolicyDefinitionTypeProperty{
//   			Name: jsii.String("name"),
//   			Values: []interface{}{
//   				&PolicyDefinitionTypeValueProperty{
//   					Value: jsii.String("value"),
//
//   					// the properties below are optional
//   					Description: jsii.String("description"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   		},
//   	},
//   	Variables: []interface{}{
//   		&PolicyDefinitionVariableProperty{
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Version: jsii.String("version"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinition.html
//
type CfnAutomatedReasoningPolicy_PolicyDefinitionProperty struct {
	// The collection of rules that define the policy logic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinition.html#cfn-bedrock-automatedreasoningpolicy-policydefinition-rules
	//
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
	// The custom types defined within the policy definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinition.html#cfn-bedrock-automatedreasoningpolicy-policydefinition-types
	//
	Types interface{} `field:"optional" json:"types" yaml:"types"`
	// The variables used within the policy definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinition.html#cfn-bedrock-automatedreasoningpolicy-policydefinition-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
	// The version of the policy definition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-automatedreasoningpolicy-policydefinition.html#cfn-bedrock-automatedreasoningpolicy-policydefinition-version
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
}

