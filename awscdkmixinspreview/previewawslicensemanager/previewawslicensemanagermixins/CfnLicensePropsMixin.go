package previewawslicensemanagermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslicensemanager/previewawslicensemanagermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a granted license.
//
// Granted licenses are licenses for products that your organization purchased from AWS Marketplace or directly from a seller who integrated their software with managed entitlements. For more information, see [Granted licenses](https://docs.aws.amazon.com/license-manager/latest/userguide/granted-licenses.html) in the *License Manager User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLicensePropsMixin := awscdkmixinspreview.Mixins.NewCfnLicensePropsMixin(&CfnLicenseMixinProps{
//   	Beneficiary: jsii.String("beneficiary"),
//   	ConsumptionConfiguration: &ConsumptionConfigurationProperty{
//   		BorrowConfiguration: &BorrowConfigurationProperty{
//   			AllowEarlyCheckIn: jsii.Boolean(false),
//   			MaxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		ProvisionalConfiguration: &ProvisionalConfigurationProperty{
//   			MaxTimeToLiveInMinutes: jsii.Number(123),
//   		},
//   		RenewType: jsii.String("renewType"),
//   	},
//   	Entitlements: []interface{}{
//   		&EntitlementProperty{
//   			AllowCheckIn: jsii.Boolean(false),
//   			MaxCount: jsii.Number(123),
//   			Name: jsii.String("name"),
//   			Overage: jsii.Boolean(false),
//   			Unit: jsii.String("unit"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	HomeRegion: jsii.String("homeRegion"),
//   	Issuer: &IssuerDataProperty{
//   		Name: jsii.String("name"),
//   		SignKey: jsii.String("signKey"),
//   	},
//   	LicenseMetadata: []interface{}{
//   		&MetadataProperty{
//   			Name: jsii.String("name"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	LicenseName: jsii.String("licenseName"),
//   	ProductName: jsii.String("productName"),
//   	ProductSku: jsii.String("productSku"),
//   	Status: jsii.String("status"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Validity: &ValidityDateFormatProperty{
//   		Begin: jsii.String("begin"),
//   		End: jsii.String("end"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-license.html
//
type CfnLicensePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnLicenseMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnLicensePropsMixin
type jsiiProxy_CfnLicensePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnLicensePropsMixin) Props() *CfnLicenseMixinProps {
	var returns *CfnLicenseMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnLicensePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::LicenseManager::License`.
func NewCfnLicensePropsMixin(props *CfnLicenseMixinProps, options *mixins.CfnPropertyMixinOptions) CfnLicensePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnLicensePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnLicensePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_licensemanager.mixins.CfnLicensePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::LicenseManager::License`.
func NewCfnLicensePropsMixin_Override(c CfnLicensePropsMixin, props *CfnLicenseMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_licensemanager.mixins.CfnLicensePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnLicensePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnLicensePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_licensemanager.mixins.CfnLicensePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnLicensePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_licensemanager.mixins.CfnLicensePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnLicensePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnLicensePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

