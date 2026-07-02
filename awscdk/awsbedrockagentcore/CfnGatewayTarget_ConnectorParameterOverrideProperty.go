package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectorParameterOverrideProperty := &ConnectorParameterOverrideProperty{
//   	Path: jsii.String("path"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Visible: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorparameteroverride.html
//
type CfnGatewayTarget_ConnectorParameterOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorparameteroverride.html#cfn-bedrockagentcore-gatewaytarget-connectorparameteroverride-path
	//
	Path *string `field:"required" json:"path" yaml:"path"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorparameteroverride.html#cfn-bedrockagentcore-gatewaytarget-connectorparameteroverride-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-connectorparameteroverride.html#cfn-bedrockagentcore-gatewaytarget-connectorparameteroverride-visible
	//
	Visible interface{} `field:"optional" json:"visible" yaml:"visible"`
}

