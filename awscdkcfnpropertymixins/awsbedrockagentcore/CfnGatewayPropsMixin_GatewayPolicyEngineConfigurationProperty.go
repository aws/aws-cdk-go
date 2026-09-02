package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gatewayPolicyEngineConfigurationProperty := &GatewayPolicyEngineConfigurationProperty{
//   	Arn: jsii.String("arn"),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.html
//
type CfnGatewayPropsMixin_GatewayPolicyEngineConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.html#cfn-bedrockagentcore-gateway-gatewaypolicyengineconfiguration-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.html#cfn-bedrockagentcore-gateway-gatewaypolicyengineconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

