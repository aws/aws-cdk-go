package awsbedrockagentcore


// The configuration for a policy engine associated with a gateway.
//
// A policy engine is a collection of policies that evaluates and authorizes agent tool calls. When associated with a gateway, the policy engine intercepts all agent requests and determines whether to allow or deny each action based on the defined policies.
//
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
	// The ARN of the policy engine.
	//
	// The policy engine contains Cedar policies that define fine-grained authorization rules specifying who can perform what actions on which resources as agents interact through the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.html#cfn-bedrockagentcore-gateway-gatewaypolicyengineconfiguration-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// The enforcement mode for the policy engine.
	//
	// LOG_ONLY - The policy engine evaluates each action against your policies and adds traces on whether tool calls would be allowed or denied, but does not enforce the decision. Use this mode to test and validate policies before enabling enforcement. ENFORCE - The policy engine evaluates actions against your policies and enforces decisions by allowing or denying agent operations. Test and validate policies in LOG_ONLY mode before enabling enforcement to avoid unintended denials or adversely affecting production traffic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewaypolicyengineconfiguration.html#cfn-bedrockagentcore-gateway-gatewaypolicyengineconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

