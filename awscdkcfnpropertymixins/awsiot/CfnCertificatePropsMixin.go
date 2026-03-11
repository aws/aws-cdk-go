package awsiot

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the `AWS::IoT::Certificate` resource to declare an AWS IoT X.509 certificate. For information about working with X.509 certificates, see [X.509 Client Certificates](https://docs.aws.amazon.com/iot/latest/developerguide/x509-client-certs.html) in the *AWS IoT Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCertificatePropsMixin := awscdkcfnpropertymixins.Aws_iot.NewCfnCertificatePropsMixin(&CfnCertificateMixinProps{
//   	CaCertificatePem: jsii.String("caCertificatePem"),
//   	CertificateMode: jsii.String("certificateMode"),
//   	CertificatePem: jsii.String("certificatePem"),
//   	CertificateSigningRequest: jsii.String("certificateSigningRequest"),
//   	Status: jsii.String("status"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html
//
type CfnCertificatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCertificateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCertificatePropsMixin
type jsiiProxy_CfnCertificatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
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

func (j *jsiiProxy_CfnCertificatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::Certificate`.
func NewCfnCertificatePropsMixin(props *CfnCertificateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCertificatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCertificatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCertificatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnCertificatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::Certificate`.
func NewCfnCertificatePropsMixin_Override(c CfnCertificatePropsMixin, props *CfnCertificateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnCertificatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCertificatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCertificatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnCertificatePropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_iot.CfnCertificatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
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

