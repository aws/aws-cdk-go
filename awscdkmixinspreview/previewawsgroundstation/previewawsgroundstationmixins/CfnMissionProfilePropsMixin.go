package previewawsgroundstationmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgroundstation/previewawsgroundstationmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Mission profiles specify parameters and provide references to config objects to define how Ground Station lists and executes contacts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMissionProfilePropsMixin := awscdkmixinspreview.Mixins.NewCfnMissionProfilePropsMixin(&CfnMissionProfileMixinProps{
//   	ContactPostPassDurationSeconds: jsii.Number(123),
//   	ContactPrePassDurationSeconds: jsii.Number(123),
//   	DataflowEdges: []interface{}{
//   		&DataflowEdgeProperty{
//   			Destination: jsii.String("destination"),
//   			Source: jsii.String("source"),
//   		},
//   	},
//   	MinimumViableContactDurationSeconds: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	StreamsKmsKey: &StreamsKmsKeyProperty{
//   		KmsAliasArn: jsii.String("kmsAliasArn"),
//   		KmsAliasName: jsii.String("kmsAliasName"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	StreamsKmsRole: jsii.String("streamsKmsRole"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TelemetrySinkConfigArn: jsii.String("telemetrySinkConfigArn"),
//   	TrackingConfigArn: jsii.String("trackingConfigArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-groundstation-missionprofile.html
//
type CfnMissionProfilePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnMissionProfileMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMissionProfilePropsMixin
type jsiiProxy_CfnMissionProfilePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnMissionProfilePropsMixin) Props() *CfnMissionProfileMixinProps {
	var returns *CfnMissionProfileMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMissionProfilePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GroundStation::MissionProfile`.
func NewCfnMissionProfilePropsMixin(props *CfnMissionProfileMixinProps, options *mixins.CfnPropertyMixinOptions) CfnMissionProfilePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMissionProfilePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMissionProfilePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnMissionProfilePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GroundStation::MissionProfile`.
func NewCfnMissionProfilePropsMixin_Override(c CfnMissionProfilePropsMixin, props *CfnMissionProfileMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnMissionProfilePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnMissionProfilePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMissionProfilePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnMissionProfilePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMissionProfilePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_groundstation.mixins.CfnMissionProfilePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMissionProfilePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMissionProfilePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

