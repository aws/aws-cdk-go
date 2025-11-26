package previewawsappsyncmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsappsync/previewawsappsyncmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes the configuration of a source API.
//
// A source API is a GraphQL API that is linked to a merged API. There can be multiple source APIs attached to each merged API. When linked to a merged API, the source API's schema, data sources, and resolvers will be combined with other linked source API data to form a new, singular API. Source APIs can originate from your account or from other accounts via Resource Access Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSourceApiAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnSourceApiAssociationPropsMixin(&CfnSourceApiAssociationMixinProps{
//   	Description: jsii.String("description"),
//   	MergedApiIdentifier: jsii.String("mergedApiIdentifier"),
//   	SourceApiAssociationConfig: &SourceApiAssociationConfigProperty{
//   		MergeType: jsii.String("mergeType"),
//   	},
//   	SourceApiIdentifier: jsii.String("sourceApiIdentifier"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-sourceapiassociation.html
//
type CfnSourceApiAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnSourceApiAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnSourceApiAssociationPropsMixin
type jsiiProxy_CfnSourceApiAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnSourceApiAssociationPropsMixin) Props() *CfnSourceApiAssociationMixinProps {
	var returns *CfnSourceApiAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnSourceApiAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AppSync::SourceApiAssociation`.
func NewCfnSourceApiAssociationPropsMixin(props *CfnSourceApiAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnSourceApiAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnSourceApiAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnSourceApiAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnSourceApiAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AppSync::SourceApiAssociation`.
func NewCfnSourceApiAssociationPropsMixin_Override(c CfnSourceApiAssociationPropsMixin, props *CfnSourceApiAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnSourceApiAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnSourceApiAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnSourceApiAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnSourceApiAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnSourceApiAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_appsync.mixins.CfnSourceApiAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnSourceApiAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnSourceApiAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

