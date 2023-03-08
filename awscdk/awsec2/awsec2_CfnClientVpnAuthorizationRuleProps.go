package awsec2


// Properties for defining a `CfnClientVpnAuthorizationRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientVpnAuthorizationRuleProps := &cfnClientVpnAuthorizationRuleProps{
//   	clientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   	targetNetworkCidr: jsii.String("targetNetworkCidr"),
//
//   	// the properties below are optional
//   	accessGroupId: jsii.String("accessGroupId"),
//   	authorizeAllGroups: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   }
//
type CfnClientVpnAuthorizationRuleProps struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string `field:"required" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
	// The IPv4 address range, in CIDR notation, of the network for which access is being authorized.
	TargetNetworkCidr *string `field:"required" json:"targetNetworkCidr" yaml:"targetNetworkCidr"`
	// The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group.
	//
	// Required if `AuthorizeAllGroups` is `false` or not specified.
	AccessGroupId *string `field:"optional" json:"accessGroupId" yaml:"accessGroupId"`
	// Indicates whether to grant access to all clients.
	//
	// Specify `true` to grant all clients who successfully establish a VPN connection access to the network. Must be set to `true` if `AccessGroupId` is not specified.
	AuthorizeAllGroups interface{} `field:"optional" json:"authorizeAllGroups" yaml:"authorizeAllGroups"`
	// A brief description of the authorization rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

