package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::Lambda::Runtime.
//
// A Lambda runtime backed by an OCI image, used by Lambda functions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnRuntimePropsMixin := awscdkcfnpropertymixins.Aws_lambda.NewCfnRuntimePropsMixin(&CfnRuntimeMixinProps{
//   	Description: jsii.String("description"),
//   	LicenseInfo: jsii.String("licenseInfo"),
//   	Name: jsii.String("name"),
//   	ReleaseLifecycle: &ReleaseLifecycleProperty{
//   		DeprecationDate: jsii.String("deprecationDate"),
//   		FunctionCreateBlockDate: jsii.String("functionCreateBlockDate"),
//   		FunctionUpdateBlockDate: jsii.String("functionUpdateBlockDate"),
//   		ReleaseState: jsii.String("releaseState"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-runtime.html
//
type CfnRuntimePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnRuntimeMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnRuntimePropsMixin
type jsiiProxy_CfnRuntimePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnRuntimePropsMixin) Props() *CfnRuntimeMixinProps {
	var returns *CfnRuntimeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuntimePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::Runtime`.
func NewCfnRuntimePropsMixin(props *CfnRuntimeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnRuntimePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnRuntimePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnRuntimePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnRuntimePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::Runtime`.
func NewCfnRuntimePropsMixin_Override(c CfnRuntimePropsMixin, props *CfnRuntimeMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnRuntimePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnRuntimePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnRuntimePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnRuntimePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRuntimePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnRuntimePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRuntimePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnRuntimePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

