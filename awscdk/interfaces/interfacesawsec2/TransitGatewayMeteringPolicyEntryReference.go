package interfacesawsec2


// A reference to a TransitGatewayMeteringPolicyEntry resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayMeteringPolicyEntryReference := &TransitGatewayMeteringPolicyEntryReference{
//   	PolicyRuleNumber: jsii.String("policyRuleNumber"),
//   	TransitGatewayMeteringPolicyId: jsii.String("transitGatewayMeteringPolicyId"),
//   }
//
type TransitGatewayMeteringPolicyEntryReference struct {
	// The PolicyRuleNumber of the TransitGatewayMeteringPolicyEntry resource.
	PolicyRuleNumber *string `field:"required" json:"policyRuleNumber" yaml:"policyRuleNumber"`
	// The TransitGatewayMeteringPolicyId of the TransitGatewayMeteringPolicyEntry resource.
	TransitGatewayMeteringPolicyId *string `field:"required" json:"transitGatewayMeteringPolicyId" yaml:"transitGatewayMeteringPolicyId"`
}

