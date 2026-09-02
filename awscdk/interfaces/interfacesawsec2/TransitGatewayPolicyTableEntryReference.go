package interfacesawsec2


// A reference to a TransitGatewayPolicyTableEntry resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayPolicyTableEntryReference := &TransitGatewayPolicyTableEntryReference{
//   	PolicyRuleNumber: jsii.String("policyRuleNumber"),
//   	TransitGatewayPolicyTableId: jsii.String("transitGatewayPolicyTableId"),
//   }
//
type TransitGatewayPolicyTableEntryReference struct {
	// The PolicyRuleNumber of the TransitGatewayPolicyTableEntry resource.
	PolicyRuleNumber *string `field:"required" json:"policyRuleNumber" yaml:"policyRuleNumber"`
	// The TransitGatewayPolicyTableId of the TransitGatewayPolicyTableEntry resource.
	TransitGatewayPolicyTableId *string `field:"required" json:"transitGatewayPolicyTableId" yaml:"transitGatewayPolicyTableId"`
}

