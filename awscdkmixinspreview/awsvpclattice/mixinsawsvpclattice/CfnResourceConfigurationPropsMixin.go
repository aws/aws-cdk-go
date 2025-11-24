package mixinsawsvpclattice

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awsvpclattice/mixinsawsvpclattice/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a resource configuration.
//
// A resource configuration defines a specific resource. You can associate a resource configuration with a service network or a VPC endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnResourceConfigurationPropsMixin(&CfnResourceConfigurationMixinProps{
//   	AllowAssociationToSharableServiceNetwork: jsii.Boolean(false),
//   	CustomDomainName: jsii.String("customDomainName"),
//   	DomainVerificationId: jsii.String("domainVerificationId"),
//   	GroupDomain: jsii.String("groupDomain"),
//   	Name: jsii.String("name"),
//   	PortRanges: []*string{
//   		jsii.String("portRanges"),
//   	},
//   	ProtocolType: jsii.String("protocolType"),
//   	ResourceConfigurationAuthType: jsii.String("resourceConfigurationAuthType"),
//   	ResourceConfigurationDefinition: &ResourceConfigurationDefinitionProperty{
//   		ArnResource: jsii.String("arnResource"),
//   		DnsResource: &DnsResourceProperty{
//   			DomainName: jsii.String("domainName"),
//   			IpAddressType: jsii.String("ipAddressType"),
//   		},
//   		IpResource: jsii.String("ipResource"),
//   	},
//   	ResourceConfigurationGroupId: jsii.String("resourceConfigurationGroupId"),
//   	ResourceConfigurationType: jsii.String("resourceConfigurationType"),
//   	ResourceGatewayId: jsii.String("resourceGatewayId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-vpclattice-resourceconfiguration.html
//
type CfnResourceConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnResourceConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourceConfigurationPropsMixin
type jsiiProxy_CfnResourceConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnResourceConfigurationPropsMixin) Props() *CfnResourceConfigurationMixinProps {
	var returns *CfnResourceConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourceConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::VpcLattice::ResourceConfiguration`.
func NewCfnResourceConfigurationPropsMixin(props *CfnResourceConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnResourceConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourceConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourceConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::VpcLattice::ResourceConfiguration`.
func NewCfnResourceConfigurationPropsMixin_Override(c CfnResourceConfigurationPropsMixin, props *CfnResourceConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnResourceConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourceConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourceConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourceConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnResourceConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

