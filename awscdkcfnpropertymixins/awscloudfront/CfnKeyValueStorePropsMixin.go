package awscloudfront

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The key value store.
//
// Use this to separate data from function code, allowing you to update data without having to publish a new version of a function. The key value store holds keys and their corresponding values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnKeyValueStorePropsMixin := awscdkcfnpropertymixins.Aws_cloudfront.NewCfnKeyValueStorePropsMixin(&CfnKeyValueStoreMixinProps{
//   	Comment: jsii.String("comment"),
//   	ImportSource: &ImportSourceProperty{
//   		SourceArn: jsii.String("sourceArn"),
//   		SourceType: jsii.String("sourceType"),
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-keyvaluestore.html
//
type CfnKeyValueStorePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnKeyValueStoreMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnKeyValueStorePropsMixin
type jsiiProxy_CfnKeyValueStorePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnKeyValueStorePropsMixin) Props() *CfnKeyValueStoreMixinProps {
	var returns *CfnKeyValueStoreMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnKeyValueStorePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::KeyValueStore`.
func NewCfnKeyValueStorePropsMixin(props *CfnKeyValueStoreMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnKeyValueStorePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnKeyValueStorePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnKeyValueStorePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyValueStorePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::KeyValueStore`.
func NewCfnKeyValueStorePropsMixin_Override(c CfnKeyValueStorePropsMixin, props *CfnKeyValueStoreMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyValueStorePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnKeyValueStorePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnKeyValueStorePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyValueStorePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnKeyValueStorePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_cloudfront.CfnKeyValueStorePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnKeyValueStorePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnKeyValueStorePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

