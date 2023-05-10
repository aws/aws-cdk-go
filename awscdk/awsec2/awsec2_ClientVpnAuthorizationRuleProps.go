package awsec2


// Properties for a ClientVpnAuthorizationRule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var clientVpnEndpoint clientVpnEndpoint
//
//   clientVpnAuthorizationRuleProps := &ClientVpnAuthorizationRuleProps{
//   	Cidr: jsii.String("cidr"),
//
//   	// the properties below are optional
//   	ClientVpnEndoint: clientVpnEndpoint,
//   	ClientVpnEndpoint: clientVpnEndpoint,
//   	Description: jsii.String("description"),
//   	GroupId: jsii.String("groupId"),
//   }
//
// Experimental.
type ClientVpnAuthorizationRuleProps struct {
	// The IPv4 address range, in CIDR notation, of the network for which access is being authorized.
	// Experimental.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// A brief description of the authorization rule.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group.
	// Experimental.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The client VPN endpoint to which to add the rule.
	// Deprecated: Use `clientVpnEndpoint` instead.
	ClientVpnEndoint IClientVpnEndpoint `field:"optional" json:"clientVpnEndoint" yaml:"clientVpnEndoint"`
	// The client VPN endpoint to which to add the rule.
	// Experimental.
	ClientVpnEndpoint IClientVpnEndpoint `field:"optional" json:"clientVpnEndpoint" yaml:"clientVpnEndpoint"`
}

