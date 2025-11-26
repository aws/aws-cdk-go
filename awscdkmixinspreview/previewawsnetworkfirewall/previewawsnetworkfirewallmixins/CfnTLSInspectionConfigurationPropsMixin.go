package previewawsnetworkfirewallmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnetworkfirewall/previewawsnetworkfirewallmixins/internal"
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
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTLSInspectionConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnTLSInspectionConfigurationPropsMixin(&CfnTLSInspectionConfigurationMixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-tlsinspectionconfiguration.html
//
type CfnTLSInspectionConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTLSInspectionConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTLSInspectionConfigurationPropsMixin
type jsiiProxy_CfnTLSInspectionConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnTLSInspectionConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkFirewall::TLSInspectionConfiguration`.
func NewCfnTLSInspectionConfigurationPropsMixin(props *CfnTLSInspectionConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTLSInspectionConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTLSInspectionConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTLSInspectionConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkFirewall::TLSInspectionConfiguration`.
func NewCfnTLSInspectionConfigurationPropsMixin_Override(c CfnTLSInspectionConfigurationPropsMixin, props *CfnTLSInspectionConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTLSInspectionConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTLSInspectionConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnTLSInspectionConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTLSInspectionConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

