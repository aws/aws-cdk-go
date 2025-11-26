package previewawsiotmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiot/previewawsiotmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a certificate provider.
//
// AWS IoT Core certificate provider lets you customize how to sign a certificate signing request (CSR) in fleet provisioning. For more information, see [Self-managed certificate signing using AWS IoT Core certificate provider](https://docs.aws.amazon.com/iot/latest/developerguide/provisioning-cert-provider.html) from the *AWS IoT Core Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCertificateProviderPropsMixin := awscdkmixinspreview.Mixins.NewCfnCertificateProviderPropsMixin(&CfnCertificateProviderMixinProps{
//   	AccountDefaultForOperations: []*string{
//   		jsii.String("accountDefaultForOperations"),
//   	},
//   	CertificateProviderName: jsii.String("certificateProviderName"),
//   	LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificateprovider.html
//
type CfnCertificateProviderPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCertificateProviderMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCertificateProviderPropsMixin
type jsiiProxy_CfnCertificateProviderPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCertificateProviderPropsMixin) Props() *CfnCertificateProviderMixinProps {
	var returns *CfnCertificateProviderMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCertificateProviderPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoT::CertificateProvider`.
func NewCfnCertificateProviderPropsMixin(props *CfnCertificateProviderMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCertificateProviderPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCertificateProviderPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCertificateProviderPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnCertificateProviderPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoT::CertificateProvider`.
func NewCfnCertificateProviderPropsMixin_Override(c CfnCertificateProviderPropsMixin, props *CfnCertificateProviderMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnCertificateProviderPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCertificateProviderPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCertificateProviderPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnCertificateProviderPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCertificateProviderPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iot.mixins.CfnCertificateProviderPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCertificateProviderPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCertificateProviderPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

