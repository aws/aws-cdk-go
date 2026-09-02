package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayPolicyEngineConfigurationProperty := &GatewayPolicyEngineConfigurationProperty{
//   	Arn: jsii.String("arn"),
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.html
//
type CfnGateway_GatewayPolicyEngineConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.html#cfn-bedrockagentcore-gateway-gatewaypolicyengineconfiguration-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.html#cfn-bedrockagentcore-gateway-gatewaypolicyengineconfiguration-mode
	//
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

