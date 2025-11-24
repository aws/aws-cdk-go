package mixinsawsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsec2/mixinsawsec2/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a Client VPN endpoint.
//
// A Client VPN endpoint is the resource you create and configure to enable and manage client VPN sessions. It is the destination endpoint at which all client VPN sessions are terminated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClientVpnEndpointPropsMixin := awscdkmixinspreview.Mixins.NewCfnClientVpnEndpointPropsMixin(&CfnClientVpnEndpointMixinProps{
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
//   	TransportProtocol: jsii.String("transportProtocol"),
//   	VpcId: jsii.String("vpcId"),
//   	VpnPort: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html
//
type CfnClientVpnEndpointPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClientVpnEndpointMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClientVpnEndpointPropsMixin
type jsiiProxy_CfnClientVpnEndpointPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnClientVpnEndpointPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::ClientVpnEndpoint`.
func NewCfnClientVpnEndpointPropsMixin(props *CfnClientVpnEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClientVpnEndpointPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClientVpnEndpointPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClientVpnEndpointPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnClientVpnEndpointPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::ClientVpnEndpoint`.
func NewCfnClientVpnEndpointPropsMixin_Override(c CfnClientVpnEndpointPropsMixin, props *CfnClientVpnEndpointMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnClientVpnEndpointPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnClientVpnEndpointPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnEndpointPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnClientVpnEndpointPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnClientVpnEndpointPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClientVpnEndpointPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
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

