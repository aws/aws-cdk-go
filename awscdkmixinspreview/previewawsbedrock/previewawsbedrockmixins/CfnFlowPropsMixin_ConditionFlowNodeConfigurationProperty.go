package previewawsbedrockmixins


// Defines a condition node in your flow.
//
// You can specify conditions that determine which node comes next in the flow. For more information, see [Node types in a flow](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-nodes.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionFlowNodeConfigurationProperty := &ConditionFlowNodeConfigurationProperty{
//   	Conditions: []interface{}{
//   		&FlowConditionProperty{
//   			Expression: jsii.String("expression"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-conditionflownodeconfiguration.html
//
type CfnFlowPropsMixin_ConditionFlowNodeConfigurationProperty struct {
	// An array of conditions.
	//
	// Each member contains the name of a condition and an expression that defines the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-conditionflownodeconfiguration.html#cfn-bedrock-flow-conditionflownodeconfiguration-conditions
	//
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

