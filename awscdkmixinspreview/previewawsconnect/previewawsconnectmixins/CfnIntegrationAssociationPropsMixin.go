package previewawsconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnect/previewawsconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies the association of an AWS resource such as Lex bot (both v1 and v2) and Lambda function with an Amazon Connect instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIntegrationAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnIntegrationAssociationPropsMixin(&CfnIntegrationAssociationMixinProps{
//   	InstanceId: jsii.String("instanceId"),
//   	IntegrationArn: jsii.String("integrationArn"),
//   	IntegrationType: jsii.String("integrationType"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-integrationassociation.html
//
type CfnIntegrationAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIntegrationAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIntegrationAssociationPropsMixin
type jsiiProxy_CfnIntegrationAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIntegrationAssociationPropsMixin) Props() *CfnIntegrationAssociationMixinProps {
	var returns *CfnIntegrationAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIntegrationAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Connect::IntegrationAssociation`.
func NewCfnIntegrationAssociationPropsMixin(props *CfnIntegrationAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIntegrationAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIntegrationAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIntegrationAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnIntegrationAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Connect::IntegrationAssociation`.
func NewCfnIntegrationAssociationPropsMixin_Override(c CfnIntegrationAssociationPropsMixin, props *CfnIntegrationAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnIntegrationAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIntegrationAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIntegrationAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnIntegrationAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIntegrationAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connect.mixins.CfnIntegrationAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIntegrationAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIntegrationAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

