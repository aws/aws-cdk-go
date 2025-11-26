package previewawsrdsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsrds/previewawsrdsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new DB shard group for Aurora Limitless Database.
//
// You must enable Aurora Limitless Database to create a DB shard group.
//
// Valid for: Aurora DB clusters only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBShardGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnDBShardGroupPropsMixin(&CfnDBShardGroupMixinProps{
//   	ComputeRedundancy: jsii.Number(123),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbShardGroupIdentifier: jsii.String("dbShardGroupIdentifier"),
//   	MaxAcu: jsii.Number(123),
//   	MinAcu: jsii.Number(123),
//   	PubliclyAccessible: jsii.Boolean(false),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbshardgroup.html
//
type CfnDBShardGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDBShardGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDBShardGroupPropsMixin
type jsiiProxy_CfnDBShardGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDBShardGroupPropsMixin) Props() *CfnDBShardGroupMixinProps {
	var returns *CfnDBShardGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDBShardGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::RDS::DBShardGroup`.
func NewCfnDBShardGroupPropsMixin(props *CfnDBShardGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDBShardGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDBShardGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDBShardGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBShardGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::RDS::DBShardGroup`.
func NewCfnDBShardGroupPropsMixin_Override(c CfnDBShardGroupPropsMixin, props *CfnDBShardGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBShardGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDBShardGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDBShardGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBShardGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDBShardGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_rds.mixins.CfnDBShardGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDBShardGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDBShardGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

