package awsbedrock


// Defines a condition node in your flow.
//
// You can specify conditions that determine which node comes next in the flow. For more information, see [Node types in Amazon Bedrock works](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-nodes.html) in the Amazon Bedrock User Guide.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionFlowNodeConfigurationProperty := &ConditionFlowNodeConfigurationProperty{
//   	Conditions: []interface{}{
//   		&FlowConditionProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			Expression: jsii.String("expression"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-conditionflownodeconfiguration.html
//
type CfnFlowVersion_ConditionFlowNodeConfigurationProperty struct {
	// An array of conditions.
	//
	// Each member contains the name of a condition and an expression that defines the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-conditionflownodeconfiguration.html#cfn-bedrock-flowversion-conditionflownodeconfiguration-conditions
	//
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
}

