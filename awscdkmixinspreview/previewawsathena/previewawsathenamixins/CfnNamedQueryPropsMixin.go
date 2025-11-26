package previewawsathenamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsathena/previewawsathenamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Athena::NamedQuery` resource specifies an Amazon Athena saved query, where `QueryString` contains the SQL query statements that make up the query.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNamedQueryPropsMixin := awscdkmixinspreview.Mixins.NewCfnNamedQueryPropsMixin(&CfnNamedQueryMixinProps{
//   	Database: jsii.String("database"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	QueryString: jsii.String("queryString"),
//   	WorkGroup: jsii.String("workGroup"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html
//
type CfnNamedQueryPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNamedQueryMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNamedQueryPropsMixin
type jsiiProxy_CfnNamedQueryPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNamedQueryPropsMixin) Props() *CfnNamedQueryMixinProps {
	var returns *CfnNamedQueryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNamedQueryPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Athena::NamedQuery`.
func NewCfnNamedQueryPropsMixin(props *CfnNamedQueryMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNamedQueryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNamedQueryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNamedQueryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_athena.mixins.CfnNamedQueryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Athena::NamedQuery`.
func NewCfnNamedQueryPropsMixin_Override(c CfnNamedQueryPropsMixin, props *CfnNamedQueryMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_athena.mixins.CfnNamedQueryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNamedQueryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNamedQueryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_athena.mixins.CfnNamedQueryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNamedQueryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_athena.mixins.CfnNamedQueryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNamedQueryPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNamedQueryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

