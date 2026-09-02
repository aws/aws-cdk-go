package awsdevopsagent

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdevopsagent/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::DevOpsAgent::Asset.
//
// An asset attached to an existing AWS DevOps Agent Space. Customer-creatable types include skill, agents_md, and attachment; call ListAssetTypes for the current authoritative set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var metadata interface{}
//
//   cfnAssetPropsMixin := awscdkcfnpropertymixins.Aws_devopsagent.NewCfnAssetPropsMixin(&CfnAssetMixinProps{
//   	AgentSpaceId: jsii.String("agentSpaceId"),
//   	AssetType: jsii.String("assetType"),
//   	Files: []interface{}{
//   		&AssetFileProperty{
//   			ContentBytes: jsii.String("contentBytes"),
//   			ContentText: jsii.String("contentText"),
//   			Metadata: metadata,
//   			Path: jsii.String("path"),
//   		},
//   	},
//   	Metadata: metadata,
//   	Zip: jsii.String("zip"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-asset.html
//
type CfnAssetPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnAssetMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAssetPropsMixin
type jsiiProxy_CfnAssetPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnAssetPropsMixin) Props() *CfnAssetMixinProps {
	var returns *CfnAssetMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAssetPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DevOpsAgent::Asset`.
func NewCfnAssetPropsMixin(props *CfnAssetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnAssetPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAssetPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAssetPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssetPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DevOpsAgent::Asset`.
func NewCfnAssetPropsMixin_Override(c CfnAssetPropsMixin, props *CfnAssetMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssetPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnAssetPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAssetPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssetPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAssetPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_devopsagent.CfnAssetPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAssetPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnAssetPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

