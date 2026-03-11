package awsmediaconvert

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediaconvert/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::MediaConvert::JobTemplate resource is an AWS Elemental MediaConvert resource type that you can use to generate transcoding jobs.
//
// When you declare this entity in your CloudFormation template, you pass in your transcoding job settings in JSON or YAML format. This settings specification must be formed in a particular way that conforms to AWS Elemental MediaConvert job validation. For more information about creating a job template model for the `SettingsJson` property, see the Remarks section later in this topic.
//
// For information about job templates, see [Working with AWS Elemental MediaConvert Job Templates](https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-job-templates.html) in the ** .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var settingsJson interface{}
//   var tags interface{}
//
//   cfnJobTemplatePropsMixin := awscdkcfnpropertymixins.Aws_mediaconvert.NewCfnJobTemplatePropsMixin(&CfnJobTemplateMixinProps{
//   	AccelerationSettings: &AccelerationSettingsProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	Category: jsii.String("category"),
//   	Description: jsii.String("description"),
//   	HopDestinations: []interface{}{
//   		&HopDestinationProperty{
//   			Priority: jsii.Number(123),
//   			Queue: jsii.String("queue"),
//   			WaitMinutes: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Priority: jsii.Number(123),
//   	Queue: jsii.String("queue"),
//   	SettingsJson: settingsJson,
//   	StatusUpdateInterval: jsii.String("statusUpdateInterval"),
//   	Tags: tags,
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconvert-jobtemplate.html
//
type CfnJobTemplatePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnJobTemplateMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnJobTemplatePropsMixin
type jsiiProxy_CfnJobTemplatePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnJobTemplatePropsMixin) Props() *CfnJobTemplateMixinProps {
	var returns *CfnJobTemplateMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnJobTemplatePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConvert::JobTemplate`.
func NewCfnJobTemplatePropsMixin(props *CfnJobTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnJobTemplatePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnJobTemplatePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnJobTemplatePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconvert.CfnJobTemplatePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConvert::JobTemplate`.
func NewCfnJobTemplatePropsMixin_Override(c CfnJobTemplatePropsMixin, props *CfnJobTemplateMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediaconvert.CfnJobTemplatePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnJobTemplatePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnJobTemplatePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediaconvert.CfnJobTemplatePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnJobTemplatePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediaconvert.CfnJobTemplatePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnJobTemplatePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnJobTemplatePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

