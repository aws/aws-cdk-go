package mixinsawsbedrock


// The complete policy definition containing rules, variables, and types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDefinitionProperty := &PolicyDefinitionProperty{
//   	Rules: []interface{}{
//   		&PolicyDefinitionRuleProperty{
//   			AlternateExpression: jsii.String("alternateExpression"),
//   			Expression: jsii.String("expression"),
//   			Id: jsii.String("id"),
//   		},
//   	},
//   	Types: []interface{}{
//   		&PolicyDefinitionTypeProperty{
//   			Description: jsii.String("description"),
//   			Name: jsii.String("name"),
//   			Values: []interface{}{
//   				&PolicyDefinitionTypeValueProperty{
//   					Description: jsii.String("description"),
//   					Value: jsii.String("value"),
//   				},
//   			},
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
type CfnAutomatedReasoningPolicyPropsMixin_PolicyDefinitionProperty struct {
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

