package previewawstransfermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawstransfer/previewawstransfermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Imports the signing and encryption certificates that you need to create local (AS2) profiles and partner profiles.
//
// You can import both the certificate and its chain in the `Certificate` parameter.
//
// After importing a certificate, AWS Transfer Family automatically creates a Amazon CloudWatch metric called `DaysUntilExpiry` that tracks the number of days until the certificate expires. The metric is based on the `InactiveDate` parameter and is published daily in the `AWS/Transfer` namespace.
//
// > It can take up to a full day after importing a certificate for Transfer Family to emit the `DaysUntilExpiry` metric to your account. > If you use the `Certificate` parameter to upload both the certificate and its chain, don't use the `CertificateChain` parameter.
//
// *CloudWatch monitoring*
//
// The `DaysUntilExpiry` metric includes the following specifications:
//
// - *Units:* Count (days)
// - *Dimensions:* `CertificateId` (always present), `Description` (if provided during certificate import)
// - *Statistics:* Minimum, Maximum, Average
// - *Frequency:* Published daily.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificatePropsMixin := awscdkmixinspreview.Mixins.NewCfnCertificatePropsMixin(&CfnCertificateMixinProps{
//   	ActiveDate: jsii.String("activeDate"),
//   	Certificate: jsii.String("certificate"),
//   	CertificateChain: jsii.String("certificateChain"),
//   	Description: jsii.String("description"),
//   	InactiveDate: jsii.String("inactiveDate"),
//   	PrivateKey: jsii.String("privateKey"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Usage: jsii.String("usage"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transfer-certificate.html
//
type CfnCertificatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCertificateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCertificatePropsMixin
type jsiiProxy_CfnCertificatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCertificatePropsMixin) Props() *CfnCertificateMixinProps {
	var returns *CfnCertificateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Transfer::Certificate`.
func NewCfnCertificatePropsMixin(props *CfnCertificateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCertificatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCertificatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCertificatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnCertificatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transfer::Certificate`.
func NewCfnCertificatePropsMixin_Override(c CfnCertificatePropsMixin, props *CfnCertificateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnCertificatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCertificatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCertificatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnCertificatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_transfer.mixins.CfnCertificatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCertificatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

