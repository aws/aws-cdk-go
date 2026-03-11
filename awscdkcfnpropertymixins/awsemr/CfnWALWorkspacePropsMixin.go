package awsemr

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-walworkspace.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWALWorkspacePropsMixin := awscdkcfnpropertymixins.Aws_emr.NewCfnWALWorkspacePropsMixin(&CfnWALWorkspaceMixinProps{
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	WalWorkspaceName: jsii.String("walWorkspaceName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-walworkspace.html
//
type CfnWALWorkspacePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWALWorkspaceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWALWorkspacePropsMixin
type jsiiProxy_CfnWALWorkspacePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWALWorkspacePropsMixin) Props() *CfnWALWorkspaceMixinProps {
	var returns *CfnWALWorkspaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWALWorkspacePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EMR::WALWorkspace`.
func NewCfnWALWorkspacePropsMixin(props *CfnWALWorkspaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWALWorkspacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWALWorkspacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWALWorkspacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnWALWorkspacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EMR::WALWorkspace`.
func NewCfnWALWorkspacePropsMixin_Override(c CfnWALWorkspacePropsMixin, props *CfnWALWorkspaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnWALWorkspacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWALWorkspacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWALWorkspacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnWALWorkspacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWALWorkspacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_emr.CfnWALWorkspacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWALWorkspacePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWALWorkspacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

