package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a Client VPN endpoint.
//
// A Client VPN endpoint is the resource you create and configure to enable and manage client VPN sessions. It is the destination endpoint at which all client VPN sessions are terminated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnClientVpnEndpointPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnClientVpnEndpointPropsMixin(&CfnClientVpnEndpointMixinProps{
//   	AuthenticationOptions: []interface{}{
//   		&ClientAuthenticationRequestProperty{
//   			ActiveDirectory: &DirectoryServiceAuthenticationRequestProperty{
//   				DirectoryId: jsii.String("directoryId"),
//   			},
//   			FederatedAuthentication: &FederatedAuthenticationRequestProperty{
//   				SamlProviderArn: jsii.String("samlProviderArn"),
//   				SelfServiceSamlProviderArn: jsii.String("selfServiceSamlProviderArn"),
//   			},
//   			MutualAuthentication: &CertificateAuthenticationRequestProperty{
//   				ClientRootCertificateChainArn: jsii.String("clientRootCertificateChainArn"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ClientCidrBlock: jsii.String("clientCidrBlock"),
//   	ClientConnectOptions: &ClientConnectOptionsProperty{
//   		Enabled: jsii.Boolean(false),
//   		LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   	},
//   	ClientLoginBannerOptions: &ClientLoginBannerOptionsProperty{
//   		BannerText: jsii.String("bannerText"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	ClientRouteEnforcementOptions: &ClientRouteEnforcementOptionsProperty{
//   		Enforced: jsii.Boolean(false),
//   	},
//   	ConnectionLogOptions: &ConnectionLogOptionsProperty{
//   		CloudwatchLogGroup: jsii.String("cloudwatchLogGroup"),
//   		CloudwatchLogStream: jsii.String("cloudwatchLogStream"),
//   		Enabled: jsii.Boolean(false),
//   	},
//   	Description: jsii.String("description"),
//   	DisconnectOnSessionTimeout: jsii.Boolean(false),
//   	DnsServers: []*string{
//   		jsii.String("dnsServers"),
//   	},
//   	EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SelfServicePortal: jsii.String("selfServicePortal"),
//   	ServerCertificateArn: jsii.String("serverCertificateArn"),
//   	SessionTimeoutHours: jsii.Number(123),
//   	SplitTunnel: jsii.Boolean(false),
//   	TagSpecifications: []interface{}{
//   		&TagSpecificationProperty{
//   			ResourceType: jsii.String("resourceType"),
//   			Tags: []CfnTag{
//   				&CfnTag{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	TrafficIpAddressType: jsii.String("trafficIpAddressType"),
//   	TransportProtocol: jsii.String("transportProtocol"),
//   	VpcId: jsii.String("vpcId"),
//   	VpnPort: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html
//
type CfnClientVpnEndpointPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnClientVpnEndpointMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClientVpnEndpointPropsMixin
type jsiiProxy_CfnClientVpnEndpointPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnClientVpnEndpointPropsMixin) Props() *CfnClientVpnEndpointMixinProps {
	var returns *CfnClientVpnEndpointMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnEndpointPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::ClientVpnEndpoint`.
func NewCfnClientVpnEndpointPropsMixin(props *CfnClientVpnEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnClientVpnEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClientVpnEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClientVpnEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnClientVpnEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::ClientVpnEndpoint`.
func NewCfnClientVpnEndpointPropsMixin_Override(c CfnClientVpnEndpointPropsMixin, props *CfnClientVpnEndpointMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnClientVpnEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnClientVpnEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnClientVpnEndpointPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClientVpnEndpointPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnClientVpnEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClientVpnEndpointPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

