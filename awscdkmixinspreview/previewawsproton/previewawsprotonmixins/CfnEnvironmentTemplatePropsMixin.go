package previewawsprotonmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsproton/previewawsprotonmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create an environment template for AWS Proton .
//
// For more information, see [Environment Templates](https://docs.aws.amazon.com/proton/latest/userguide/ag-templates.html) in the *AWS Proton User Guide* .
//
// You can create an environment template in one of the two following ways:
//
// - Register and publish a *standard* environment template that instructs AWS Proton to deploy and manage environment infrastructure.
// - Register and publish a *customer managed* environment template that connects AWS Proton to your existing provisioned infrastructure that you manage. AWS Proton *doesn't* manage your existing provisioned infrastructure. To create an environment template for customer provisioned and managed infrastructure, include the `provisioning` parameter and set the value to `CUSTOMER_MANAGED` . For more information, see [Register and publish an environment template](https://docs.aws.amazon.com/proton/latest/userguide/template-create.html) in the *AWS Proton User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEnvironmentTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnEnvironmentTemplatePropsMixin(&CfnEnvironmentTemplateMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	Name: jsii.String("name"),
//   	Provisioning: jsii.String("provisioning"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-environmenttemplate.html
//
type CfnEnvironmentTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnEnvironmentTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnEnvironmentTemplatePropsMixin
type jsiiProxy_CfnEnvironmentTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnEnvironmentTemplatePropsMixin) Props() *CfnEnvironmentTemplateMixinProps {
	var returns *CfnEnvironmentTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnEnvironmentTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Proton::EnvironmentTemplate`.
func NewCfnEnvironmentTemplatePropsMixin(props *CfnEnvironmentTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnEnvironmentTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnEnvironmentTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnEnvironmentTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnEnvironmentTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Proton::EnvironmentTemplate`.
func NewCfnEnvironmentTemplatePropsMixin_Override(c CfnEnvironmentTemplatePropsMixin, props *CfnEnvironmentTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnEnvironmentTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnEnvironmentTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnEnvironmentTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnEnvironmentTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnEnvironmentTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnEnvironmentTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnEnvironmentTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnEnvironmentTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

