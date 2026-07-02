package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var parameterValues interface{}
//
//   connectorConfigurationProperty := &ConnectorConfigurationProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	ParameterOverrides: []interface{}{
//   		&ConnectorParameterOverrideProperty{
//   			Path: jsii.String("path"),
//
//   			// the properties below are optional
//   			Description: jsii.String("description"),
//   			Visible: jsii.Boolean(false),
//   		},
//   	},
//   	ParameterValues: parameterValues,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorconfiguration.html
//
type CfnGatewayTarget_ConnectorConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectorconfiguration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectorconfiguration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectorconfiguration-parameteroverrides
	//
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorconfiguration.html#cfn-bedrockagentcore-gatewaytarget-connectorconfiguration-parametervalues
	//
	ParameterValues interface{} `field:"optional" json:"parameterValues" yaml:"parameterValues"`
}

