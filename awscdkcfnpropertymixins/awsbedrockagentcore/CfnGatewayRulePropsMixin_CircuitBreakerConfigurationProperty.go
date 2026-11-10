package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   circuitBreakerConfigurationProperty := &CircuitBreakerConfigurationProperty{
//   	Mode: jsii.String("mode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-circuitbreakerconfiguration.html
//
type CfnGatewayRulePropsMixin_CircuitBreakerConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewayrule-circuitbreakerconfiguration.html#cfn-bedrockagentcore-gatewayrule-circuitbreakerconfiguration-mode
	//
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

