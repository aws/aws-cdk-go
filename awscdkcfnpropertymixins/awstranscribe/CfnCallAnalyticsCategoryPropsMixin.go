package awstranscribe

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awstranscribe/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource type definition for AWS::Transcribe::CallAnalyticsCategory.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnCallAnalyticsCategoryPropsMixin := awscdkcfnpropertymixins.Aws_transcribe.NewCfnCallAnalyticsCategoryPropsMixin(&CfnCallAnalyticsCategoryMixinProps{
//   	CategoryName: jsii.String("categoryName"),
//   	InputType: jsii.String("inputType"),
//   	Rules: []interface{}{
//   		&RuleProperty{
//   			InterruptionFilter: &InterruptionFilterProperty{
//   				AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   					EndTime: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartTime: jsii.Number(123),
//   				},
//   				Negate: jsii.Boolean(false),
//   				ParticipantRole: jsii.String("participantRole"),
//   				RelativeTimeRange: &RelativeTimeRangeProperty{
//   					EndPercentage: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartPercentage: jsii.Number(123),
//   				},
//   				Threshold: jsii.Number(123),
//   			},
//   			NonTalkTimeFilter: &NonTalkTimeFilterProperty{
//   				AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   					EndTime: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartTime: jsii.Number(123),
//   				},
//   				Negate: jsii.Boolean(false),
//   				RelativeTimeRange: &RelativeTimeRangeProperty{
//   					EndPercentage: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartPercentage: jsii.Number(123),
//   				},
//   				Threshold: jsii.Number(123),
//   			},
//   			SentimentFilter: &SentimentFilterProperty{
//   				AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   					EndTime: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartTime: jsii.Number(123),
//   				},
//   				Negate: jsii.Boolean(false),
//   				ParticipantRole: jsii.String("participantRole"),
//   				RelativeTimeRange: &RelativeTimeRangeProperty{
//   					EndPercentage: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartPercentage: jsii.Number(123),
//   				},
//   				Sentiments: []*string{
//   					jsii.String("sentiments"),
//   				},
//   			},
//   			TranscriptFilter: &TranscriptFilterProperty{
//   				AbsoluteTimeRange: &AbsoluteTimeRangeProperty{
//   					EndTime: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartTime: jsii.Number(123),
//   				},
//   				Negate: jsii.Boolean(false),
//   				ParticipantRole: jsii.String("participantRole"),
//   				RelativeTimeRange: &RelativeTimeRangeProperty{
//   					EndPercentage: jsii.Number(123),
//   					First: jsii.Number(123),
//   					Last: jsii.Number(123),
//   					StartPercentage: jsii.Number(123),
//   				},
//   				Targets: []*string{
//   					jsii.String("targets"),
//   				},
//   				TranscriptFilterType: jsii.String("transcriptFilterType"),
//   			},
//   		},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-transcribe-callanalyticscategory.html
//
type CfnCallAnalyticsCategoryPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnCallAnalyticsCategoryMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCallAnalyticsCategoryPropsMixin
type jsiiProxy_CfnCallAnalyticsCategoryPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnCallAnalyticsCategoryPropsMixin) Props() *CfnCallAnalyticsCategoryMixinProps {
	var returns *CfnCallAnalyticsCategoryMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCallAnalyticsCategoryPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Transcribe::CallAnalyticsCategory`.
func NewCfnCallAnalyticsCategoryPropsMixin(props *CfnCallAnalyticsCategoryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnCallAnalyticsCategoryPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCallAnalyticsCategoryPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCallAnalyticsCategoryPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Transcribe::CallAnalyticsCategory`.
func NewCfnCallAnalyticsCategoryPropsMixin_Override(c CfnCallAnalyticsCategoryPropsMixin, props *CfnCallAnalyticsCategoryMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnCallAnalyticsCategoryPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCallAnalyticsCategoryPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCallAnalyticsCategoryPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCallAnalyticsCategoryPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnCallAnalyticsCategoryPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

