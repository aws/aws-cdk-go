package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A resource discovery is an IPAM component that enables IPAM to manage and monitor resources that belong to the owning account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIPAMResourceDiscoveryPropsMixin := awscdkmixinspreview.Mixins.NewCfnIPAMResourceDiscoveryPropsMixin(&CfnIPAMResourceDiscoveryMixinProps{
//   	Description: jsii.String("description"),
//   	OperatingRegions: []interface{}{
//   		&IpamOperatingRegionProperty{
//   			RegionName: jsii.String("regionName"),
//   		},
//   	},
//   	OrganizationalUnitExclusions: []interface{}{
//   		&IpamResourceDiscoveryOrganizationalUnitExclusionProperty{
//   			OrganizationsEntityPath: jsii.String("organizationsEntityPath"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipamresourcediscovery.html
//
type CfnIPAMResourceDiscoveryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIPAMResourceDiscoveryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIPAMResourceDiscoveryPropsMixin
type jsiiProxy_CfnIPAMResourceDiscoveryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIPAMResourceDiscoveryPropsMixin) Props() *CfnIPAMResourceDiscoveryMixinProps {
	var returns *CfnIPAMResourceDiscoveryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMResourceDiscoveryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::IPAMResourceDiscovery`.
func NewCfnIPAMResourceDiscoveryPropsMixin(props *CfnIPAMResourceDiscoveryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIPAMResourceDiscoveryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIPAMResourceDiscoveryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIPAMResourceDiscoveryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnIPAMResourceDiscoveryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::IPAMResourceDiscovery`.
func NewCfnIPAMResourceDiscoveryPropsMixin_Override(c CfnIPAMResourceDiscoveryPropsMixin, props *CfnIPAMResourceDiscoveryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnIPAMResourceDiscoveryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIPAMResourceDiscoveryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIPAMResourceDiscoveryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnIPAMResourceDiscoveryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIPAMResourceDiscoveryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnIPAMResourceDiscoveryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIPAMResourceDiscoveryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIPAMResourceDiscoveryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

