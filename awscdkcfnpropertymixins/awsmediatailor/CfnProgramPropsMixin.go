package awsmediatailor

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsmediatailor/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource schema for AWS::MediaTailor::Program.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnProgramPropsMixin := awscdkcfnpropertymixins.Aws_mediatailor.NewCfnProgramPropsMixin(&CfnProgramMixinProps{
//   	AdBreaks: []interface{}{
//   		&AdBreakProperty{
//   			AdBreakMetadata: []interface{}{
//   				&KeyValuePairProperty{
//   					Key: jsii.String("key"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			MessageType: jsii.String("messageType"),
//   			OffsetMillis: jsii.Number(123),
//   			Slate: &SlateSourceProperty{
//   				SourceLocationName: jsii.String("sourceLocationName"),
//   				VodSourceName: jsii.String("vodSourceName"),
//   			},
//   			SpliceInsertMessage: &SpliceInsertMessageProperty{
//   				AvailNum: jsii.Number(123),
//   				AvailsExpected: jsii.Number(123),
//   				SpliceEventId: jsii.Number(123),
//   				UniqueProgramId: jsii.Number(123),
//   			},
//   			TimeSignalMessage: &TimeSignalMessageProperty{
//   				SegmentationDescriptors: []interface{}{
//   					&SegmentationDescriptorProperty{
//   						SegmentationEventId: jsii.Number(123),
//   						SegmentationTypeId: jsii.Number(123),
//   						SegmentationUpid: jsii.String("segmentationUpid"),
//   						SegmentationUpidType: jsii.Number(123),
//   						SegmentNum: jsii.Number(123),
//   						SegmentsExpected: jsii.Number(123),
//   						SubSegmentNum: jsii.Number(123),
//   						SubSegmentsExpected: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	AudienceMedia: []interface{}{
//   		&AudienceMediaProperty{
//   			AlternateMedia: []interface{}{
//   				&AlternateMediaProperty{
//   					AdBreaks: []interface{}{
//   						&AdBreakProperty{
//   							AdBreakMetadata: []interface{}{
//   								&KeyValuePairProperty{
//   									Key: jsii.String("key"),
//   									Value: jsii.String("value"),
//   								},
//   							},
//   							MessageType: jsii.String("messageType"),
//   							OffsetMillis: jsii.Number(123),
//   							Slate: &SlateSourceProperty{
//   								SourceLocationName: jsii.String("sourceLocationName"),
//   								VodSourceName: jsii.String("vodSourceName"),
//   							},
//   							SpliceInsertMessage: &SpliceInsertMessageProperty{
//   								AvailNum: jsii.Number(123),
//   								AvailsExpected: jsii.Number(123),
//   								SpliceEventId: jsii.Number(123),
//   								UniqueProgramId: jsii.Number(123),
//   							},
//   							TimeSignalMessage: &TimeSignalMessageProperty{
//   								SegmentationDescriptors: []interface{}{
//   									&SegmentationDescriptorProperty{
//   										SegmentationEventId: jsii.Number(123),
//   										SegmentationTypeId: jsii.Number(123),
//   										SegmentationUpid: jsii.String("segmentationUpid"),
//   										SegmentationUpidType: jsii.Number(123),
//   										SegmentNum: jsii.Number(123),
//   										SegmentsExpected: jsii.Number(123),
//   										SubSegmentNum: jsii.Number(123),
//   										SubSegmentsExpected: jsii.Number(123),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					ClipRange: &ClipRangeProperty{
//   						EndOffsetMillis: jsii.Number(123),
//   						StartOffsetMillis: jsii.Number(123),
//   					},
//   					DurationMillis: jsii.Number(123),
//   					LiveSourceName: jsii.String("liveSourceName"),
//   					ScheduledStartTimeMillis: jsii.Number(123),
//   					SourceLocationName: jsii.String("sourceLocationName"),
//   					VodSourceName: jsii.String("vodSourceName"),
//   				},
//   			},
//   			Audience: jsii.String("audience"),
//   		},
//   	},
//   	ChannelName: jsii.String("channelName"),
//   	LiveSourceName: jsii.String("liveSourceName"),
//   	ProgramName: jsii.String("programName"),
//   	ScheduleConfiguration: &ScheduleConfigurationProperty{
//   		ClipRange: &ClipRangeProperty{
//   			EndOffsetMillis: jsii.Number(123),
//   			StartOffsetMillis: jsii.Number(123),
//   		},
//   		Transition: &TransitionProperty{
//   			DurationMillis: jsii.Number(123),
//   			RelativePosition: jsii.String("relativePosition"),
//   			RelativeProgram: jsii.String("relativeProgram"),
//   			ScheduledStartTimeMillis: jsii.Number(123),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SourceLocationName: jsii.String("sourceLocationName"),
//   	VodSourceName: jsii.String("vodSourceName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediatailor-program.html
//
type CfnProgramPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnProgramMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProgramPropsMixin
type jsiiProxy_CfnProgramPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnProgramPropsMixin) Props() *CfnProgramMixinProps {
	var returns *CfnProgramMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProgramPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaTailor::Program`.
func NewCfnProgramPropsMixin(props *CfnProgramMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnProgramPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProgramPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProgramPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnProgramPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaTailor::Program`.
func NewCfnProgramPropsMixin_Override(c CfnProgramPropsMixin, props *CfnProgramMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnProgramPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnProgramPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProgramPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnProgramPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProgramPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_mediatailor.CfnProgramPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProgramPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnProgramPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

