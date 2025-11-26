package previewawsprotonmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsproton/previewawsprotonmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Create a service template.
//
// The administrator creates a service template to define standardized infrastructure and an optional CI/CD service pipeline. Developers, in turn, select the service template from AWS Proton . If the selected service template includes a service pipeline definition, they provide a link to their source code repository. AWS Proton then deploys and manages the infrastructure defined by the selected service template. For more information, see [AWS Proton templates](https://docs.aws.amazon.com/proton/latest/userguide/ag-templates.html) in the *AWS Proton User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnServiceTemplatePropsMixin := awscdkmixinspreview.Mixins.NewCfnServiceTemplatePropsMixin(&CfnServiceTemplateMixinProps{
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	EncryptionKey: jsii.String("encryptionKey"),
//   	Name: jsii.String("name"),
//   	PipelineProvisioning: jsii.String("pipelineProvisioning"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-proton-servicetemplate.html
//
type CfnServiceTemplatePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServiceTemplateMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServiceTemplatePropsMixin
type jsiiProxy_CfnServiceTemplatePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServiceTemplatePropsMixin) Props() *CfnServiceTemplateMixinProps {
	var returns *CfnServiceTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServiceTemplatePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Proton::ServiceTemplate`.
func NewCfnServiceTemplatePropsMixin(props *CfnServiceTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServiceTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServiceTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServiceTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnServiceTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Proton::ServiceTemplate`.
func NewCfnServiceTemplatePropsMixin_Override(c CfnServiceTemplatePropsMixin, props *CfnServiceTemplateMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnServiceTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServiceTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServiceTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnServiceTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServiceTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_proton.mixins.CfnServiceTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServiceTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnServiceTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

