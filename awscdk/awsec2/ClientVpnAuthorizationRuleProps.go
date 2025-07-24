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
//   	ClientVpnEndpoint: clientVpnEndpoint,
//   	Description: jsii.String("description"),
//   	GroupId: jsii.String("groupId"),
//   }
//
type ClientVpnAuthorizationRuleProps struct {
	// The IPv4 address range, in CIDR notation, of the network for which access is being authorized.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// A brief description of the authorization rule.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group.
	// Default: - authorize all groups.
	//
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The client VPN endpoint to which to add the rule.
	// Default: clientVpnEndpoint is required.
	//
	ClientVpnEndpoint IClientVpnEndpoint `field:"optional" json:"clientVpnEndpoint" yaml:"clientVpnEndpoint"`
}

