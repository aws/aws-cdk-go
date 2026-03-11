package awsnetworkfirewall

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The object that defines a TLS inspection configuration.
//
// AWS Network Firewall uses a TLS inspection configuration to decrypt traffic. Network Firewall re-encrypts the traffic before sending it to its destination.
//
// To use a TLS inspection configuration, you add it to a new Network Firewall firewall policy, then you apply the firewall policy to a firewall. Network Firewall acts as a proxy service to decrypt and inspect the traffic traveling through your firewalls. You can reference a TLS inspection configuration from more than one firewall policy, and you can use a firewall policy in more than one firewall. For more information about using TLS inspection configurations, see [Inspecting SSL/TLS traffic with TLS inspection configurations](https://docs.aws.amazon.com/network-firewall/latest/developerguide/tls-inspection.html) in the *AWS Network Firewall Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTLSInspectionConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_networkfirewall.NewCfnTLSInspectionConfigurationPropsMixin(&CfnTLSInspectionConfigurationMixinProps{
//   	Description: jsii.String("description"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TlsInspectionConfiguration: &TLSInspectionConfigurationProperty{
//   		ServerCertificateConfigurations: []interface{}{
//   			&ServerCertificateConfigurationProperty{
//   				CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   				CheckCertificateRevocationStatus: &CheckCertificateRevocationStatusProperty{
//   					RevokedStatusAction: jsii.String("revokedStatusAction"),
//   					UnknownStatusAction: jsii.String("unknownStatusAction"),
//   				},
//   				Scopes: []interface{}{
//   					&ServerCertificateScopeProperty{
//   						DestinationPorts: []interface{}{
//   							&PortRangeProperty{
//   								FromPort: jsii.Number(123),
//   								ToPort: jsii.Number(123),
//   							},
//   						},
//   						Destinations: []interface{}{
//   							&AddressProperty{
//   								AddressDefinition: jsii.String("addressDefinition"),
//   							},
//   						},
//   						Protocols: []interface{}{
//   							jsii.Number(123),
//   						},
//   						SourcePorts: []interface{}{
//   							&PortRangeProperty{
//   								FromPort: jsii.Number(123),
//   								ToPort: jsii.Number(123),
//   							},
//   						},
//   						Sources: []interface{}{
//   							&AddressProperty{
//   								AddressDefinition: jsii.String("addressDefinition"),
//   							},
//   						},
//   					},
//   				},
//   				ServerCertificates: []interface{}{
//   					&ServerCertificateProperty{
//   						ResourceArn: jsii.String("resourceArn"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	TlsInspectionConfigurationName: jsii.String("tlsInspectionConfigurationName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html
//
type CfnTLSInspectionConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTLSInspectionConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTLSInspectionConfigurationPropsMixin
type jsiiProxy_CfnTLSInspectionConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTLSInspectionConfigurationPropsMixin) Props() *CfnTLSInspectionConfigurationMixinProps {
	var returns *CfnTLSInspectionConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTLSInspectionConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkFirewall::TLSInspectionConfiguration`.
func NewCfnTLSInspectionConfigurationPropsMixin(props *CfnTLSInspectionConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTLSInspectionConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTLSInspectionConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTLSInspectionConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkFirewall::TLSInspectionConfiguration`.
func NewCfnTLSInspectionConfigurationPropsMixin_Override(c CfnTLSInspectionConfigurationPropsMixin, props *CfnTLSInspectionConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTLSInspectionConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTLSInspectionConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTLSInspectionConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_networkfirewall.CfnTLSInspectionConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTLSInspectionConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTLSInspectionConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

