package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// User-based authentication for a client VPN endpoint.
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
// Experimental.
type ClientVpnUserBasedAuthentication interface {
	// Renders the user based authentication.
	// Experimental.
	Render() interface{}
}

// The jsii proxy struct for ClientVpnUserBasedAuthentication
type jsiiProxy_ClientVpnUserBasedAuthentication struct {
	_ byte // padding
}

// Experimental.
func NewClientVpnUserBasedAuthentication_Override(c ClientVpnUserBasedAuthentication) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_ec2.ClientVpnUserBasedAuthentication",
		nil, // no parameters
		c,
	)
}

// Active Directory authentication.
// Experimental.
func ClientVpnUserBasedAuthentication_ActiveDirectory(directoryId *string) ClientVpnUserBasedAuthentication {
	_init_.Initialize()

	if err := validateClientVpnUserBasedAuthentication_ActiveDirectoryParameters(directoryId); err != nil {
		panic(err)
	}
	var returns ClientVpnUserBasedAuthentication

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.ClientVpnUserBasedAuthentication",
		"activeDirectory",
		[]interface{}{directoryId},
		&returns,
	)

	return returns
}

// Federated authentication.
// Experimental.
func ClientVpnUserBasedAuthentication_Federated(samlProvider awsiam.ISamlProvider, selfServiceSamlProvider awsiam.ISamlProvider) ClientVpnUserBasedAuthentication {
	_init_.Initialize()

	if err := validateClientVpnUserBasedAuthentication_FederatedParameters(samlProvider); err != nil {
		panic(err)
	}
	var returns ClientVpnUserBasedAuthentication

	_jsii_.StaticInvoke(
		"monocdk.aws_ec2.ClientVpnUserBasedAuthentication",
		"federated",
		[]interface{}{samlProvider, selfServiceSamlProvider},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClientVpnUserBasedAuthentication) Render() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"render",
		nil, // no parameters
		&returns,
	)

	return returns
}

