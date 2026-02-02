package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// In IPAM, a pool is a collection of contiguous IP addresses CIDRs.
//
// Pools enable you to organize your IP addresses according to your routing and security needs. For example, if you have separate routing and security needs for development and production applications, you can create a pool for each.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIPAMPoolPropsMixin := awscdkmixinspreview.Mixins.NewCfnIPAMPoolPropsMixin(&CfnIPAMPoolMixinProps{
//   	AddressFamily: jsii.String("addressFamily"),
//   	AllocationDefaultNetmaskLength: jsii.Number(123),
//   	AllocationMaxNetmaskLength: jsii.Number(123),
//   	AllocationMinNetmaskLength: jsii.Number(123),
//   	AllocationResourceTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	AutoImport: jsii.Boolean(false),
//   	AwsService: jsii.String("awsService"),
//   	Description: jsii.String("description"),
//   	IpamScopeId: jsii.String("ipamScopeId"),
//   	Locale: jsii.String("locale"),
//   	ProvisionedCidrs: []interface{}{
//   		&ProvisionedCidrProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	PublicIpSource: jsii.String("publicIpSource"),
//   	PubliclyAdvertisable: jsii.Boolean(false),
//   	SourceIpamPoolId: jsii.String("sourceIpamPoolId"),
//   	SourceResource: &SourceResourceProperty{
//   		ResourceId: jsii.String("resourceId"),
//   		ResourceOwner: jsii.String("resourceOwner"),
//   		ResourceRegion: jsii.String("resourceRegion"),
//   		ResourceType: jsii.String("resourceType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html
//
type CfnIPAMPoolPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIPAMPoolMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIPAMPoolPropsMixin
type jsiiProxy_CfnIPAMPoolPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIPAMPoolPropsMixin) Props() *CfnIPAMPoolMixinProps {
	var returns *CfnIPAMPoolMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIPAMPoolPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::IPAMPool`.
func NewCfnIPAMPoolPropsMixin(props *CfnIPAMPoolMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIPAMPoolPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIPAMPoolPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIPAMPoolPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnIPAMPoolPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::IPAMPool`.
func NewCfnIPAMPoolPropsMixin_Override(c CfnIPAMPoolPropsMixin, props *CfnIPAMPoolMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnIPAMPoolPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIPAMPoolPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIPAMPoolPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnIPAMPoolPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIPAMPoolPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnIPAMPoolPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIPAMPoolPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnIPAMPoolPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

