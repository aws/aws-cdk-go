package previewawsivsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsivs/previewawsivsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::IVS::IngestConfiguration` resource specifies an ingest protocol to be used for a stage.
//
// For more information, see [Stream Ingest](https://docs.aws.amazon.com/ivs/latest/RealTimeUserGuide/rt-stream-ingest.html) in the *Amazon IVS Real-Time Streaming User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIngestConfigurationPropsMixin := awscdkmixinspreview.Mixins.NewCfnIngestConfigurationPropsMixin(&CfnIngestConfigurationMixinProps{
//   	IngestProtocol: jsii.String("ingestProtocol"),
//   	InsecureIngest: jsii.Boolean(false),
//   	Name: jsii.String("name"),
//   	StageArn: jsii.String("stageArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserId: jsii.String("userId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ivs-ingestconfiguration.html
//
type CfnIngestConfigurationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnIngestConfigurationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnIngestConfigurationPropsMixin
type jsiiProxy_CfnIngestConfigurationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnIngestConfigurationPropsMixin) Props() *CfnIngestConfigurationMixinProps {
	var returns *CfnIngestConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnIngestConfigurationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IVS::IngestConfiguration`.
func NewCfnIngestConfigurationPropsMixin(props *CfnIngestConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnIngestConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnIngestConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnIngestConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnIngestConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IVS::IngestConfiguration`.
func NewCfnIngestConfigurationPropsMixin_Override(c CfnIngestConfigurationPropsMixin, props *CfnIngestConfigurationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnIngestConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnIngestConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnIngestConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnIngestConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnIngestConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ivs.mixins.CfnIngestConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnIngestConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnIngestConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

