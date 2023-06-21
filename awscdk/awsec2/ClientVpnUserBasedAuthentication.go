package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
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
type ClientVpnUserBasedAuthentication interface {
	// Renders the user based authentication.
	Render() interface{}
}

// The jsii proxy struct for ClientVpnUserBasedAuthentication
type jsiiProxy_ClientVpnUserBasedAuthentication struct {
	_ byte // padding
}

func NewClientVpnUserBasedAuthentication_Override(c ClientVpnUserBasedAuthentication) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.ClientVpnUserBasedAuthentication",
		nil, // no parameters
		c,
	)
}

// Active Directory authentication.
func ClientVpnUserBasedAuthentication_ActiveDirectory(directoryId *string) ClientVpnUserBasedAuthentication {
	_init_.Initialize()

	if err := validateClientVpnUserBasedAuthentication_ActiveDirectoryParameters(directoryId); err != nil {
		panic(err)
	}
	var returns ClientVpnUserBasedAuthentication

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.ClientVpnUserBasedAuthentication",
		"activeDirectory",
		[]interface{}{directoryId},
		&returns,
	)

	return returns
}

// Federated authentication.
func ClientVpnUserBasedAuthentication_Federated(samlProvider awsiam.ISamlProvider, selfServiceSamlProvider awsiam.ISamlProvider) ClientVpnUserBasedAuthentication {
	_init_.Initialize()

	if err := validateClientVpnUserBasedAuthentication_FederatedParameters(samlProvider); err != nil {
		panic(err)
	}
	var returns ClientVpnUserBasedAuthentication

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.ClientVpnUserBasedAuthentication",
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

