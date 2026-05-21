package awsawsexternalanthropic

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsawsexternalanthropic/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::AWSExternalAnthropic::Workspace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnWorkspacePropsMixin := awscdkcfnpropertymixins.Aws_awsexternalanthropic.NewCfnWorkspacePropsMixin(&CfnWorkspaceMixinProps{
//   	DataResidency: &DataResidencyProperty{
//   		AllowedInferenceGeos: []*string{
//   			jsii.String("allowedInferenceGeos"),
//   		},
//   		DefaultInferenceGeo: jsii.String("defaultInferenceGeo"),
//   		WorkspaceGeo: jsii.String("workspaceGeo"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-awsexternalanthropic-workspace.html
//
type CfnWorkspacePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnWorkspaceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnWorkspacePropsMixin
type jsiiProxy_CfnWorkspacePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnWorkspacePropsMixin) Props() *CfnWorkspaceMixinProps {
	var returns *CfnWorkspaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnWorkspacePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::AWSExternalAnthropic::Workspace`.
func NewCfnWorkspacePropsMixin(props *CfnWorkspaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnWorkspacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnWorkspacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnWorkspacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_awsexternalanthropic.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::AWSExternalAnthropic::Workspace`.
func NewCfnWorkspacePropsMixin_Override(c CfnWorkspacePropsMixin, props *CfnWorkspaceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_awsexternalanthropic.CfnWorkspacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnWorkspacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnWorkspacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_awsexternalanthropic.CfnWorkspacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnWorkspacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_awsexternalanthropic.CfnWorkspacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnWorkspacePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnWorkspacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

