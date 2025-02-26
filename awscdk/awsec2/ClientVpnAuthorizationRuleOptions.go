package awsec2


// Options for a ClientVpnAuthorizationRule.
//
// Example:
//   endpoint := vpc.addClientVpnEndpoint(jsii.String("Endpoint"), &ClientVpnEndpointOptions{
//   	Cidr: jsii.String("10.100.0.0/16"),
//   	ServerCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/server-certificate-id"),
//   	UserBasedAuthentication: ec2.ClientVpnUserBasedAuthentication_Federated(samlProvider),
//   	AuthorizeAllUsersToVpcCidr: jsii.Boolean(false),
//   })
//
//   endpoint.AddAuthorizationRule(jsii.String("Rule"), &ClientVpnAuthorizationRuleOptions{
//   	Cidr: jsii.String("10.0.10.0/32"),
//   	GroupId: jsii.String("group-id"),
//   })
//
type ClientVpnAuthorizationRuleOptions struct {
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
}

