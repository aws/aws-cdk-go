package awsoutposts

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsoutposts/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Definition of AWS::Outposts::Site Resource Type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnSitePropsMixin := awscdkcfnpropertymixins.Aws_outposts.NewCfnSitePropsMixin(&CfnSiteMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Notes: jsii.String("notes"),
//   	OperatingAddress: &AddressProperty{
//   		AddressLine1: jsii.String("addressLine1"),
//   		AddressLine2: jsii.String("addressLine2"),
//   		AddressLine3: jsii.String("addressLine3"),
//   		City: jsii.String("city"),
//   		ContactName: jsii.String("contactName"),
//   		ContactPhoneNumber: jsii.String("contactPhoneNumber"),
//   		CountryCode: jsii.String("countryCode"),
//   		DistrictOrCounty: jsii.String("districtOrCounty"),
//   		Municipality: jsii.String("municipality"),
//   		PostalCode: jsii.String("postalCode"),
//   		StateOrRegion: jsii.String("stateOrRegion"),
//   	},
//   	RackPhysicalProperties: &RackPhysicalPropertiesProperty{
//   		FiberOpticCableType: jsii.String("fiberOpticCableType"),
//   		MaximumSupportedWeightLbs: jsii.String("maximumSupportedWeightLbs"),
//   		OpticalStandard: jsii.String("opticalStandard"),
//   		PowerConnector: jsii.String("powerConnector"),
//   		PowerDrawKva: jsii.String("powerDrawKva"),
//   		PowerFeedDrop: jsii.String("powerFeedDrop"),
//   		PowerPhase: jsii.String("powerPhase"),
//   		UplinkCount: jsii.String("uplinkCount"),
//   		UplinkGbps: jsii.String("uplinkGbps"),
//   	},
//   	ShippingAddress: &AddressProperty{
//   		AddressLine1: jsii.String("addressLine1"),
//   		AddressLine2: jsii.String("addressLine2"),
//   		AddressLine3: jsii.String("addressLine3"),
//   		City: jsii.String("city"),
//   		ContactName: jsii.String("contactName"),
//   		ContactPhoneNumber: jsii.String("contactPhoneNumber"),
//   		CountryCode: jsii.String("countryCode"),
//   		DistrictOrCounty: jsii.String("districtOrCounty"),
//   		Municipality: jsii.String("municipality"),
//   		PostalCode: jsii.String("postalCode"),
//   		StateOrRegion: jsii.String("stateOrRegion"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-outposts-site.html
//
type CfnSitePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnSiteMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSitePropsMixin
type jsiiProxy_CfnSitePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnSitePropsMixin) Props() *CfnSiteMixinProps {
	var returns *CfnSiteMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSitePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Outposts::Site`.
func NewCfnSitePropsMixin(props *CfnSiteMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnSitePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSitePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSitePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_outposts.CfnSitePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Outposts::Site`.
func NewCfnSitePropsMixin_Override(c CfnSitePropsMixin, props *CfnSiteMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_outposts.CfnSitePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnSitePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSitePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_outposts.CfnSitePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSitePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_outposts.CfnSitePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSitePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnSitePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

