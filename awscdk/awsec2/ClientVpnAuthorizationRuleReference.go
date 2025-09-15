package awsec2


// A reference to a ClientVpnAuthorizationRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientVpnAuthorizationRuleReference := &ClientVpnAuthorizationRuleReference{
//   	ClientVpnAuthorizationRuleId: jsii.String("clientVpnAuthorizationRuleId"),
//   }
//
type ClientVpnAuthorizationRuleReference struct {
	// The Id of the ClientVpnAuthorizationRule resource.
	ClientVpnAuthorizationRuleId *string `field:"required" json:"clientVpnAuthorizationRuleId" yaml:"clientVpnAuthorizationRuleId"`
}

