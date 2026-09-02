package interfacesawsbedrockagentcore


// A reference to a GatewayRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRuleReference := &GatewayRuleReference{
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	RuleId: jsii.String("ruleId"),
//   }
//
type GatewayRuleReference struct {
	// The GatewayIdentifier of the GatewayRule resource.
	GatewayIdentifier *string `field:"required" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
	// The RuleId of the GatewayRule resource.
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
}

