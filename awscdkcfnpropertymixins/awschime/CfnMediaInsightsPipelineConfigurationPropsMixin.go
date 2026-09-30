package awschime

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awschime/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for an Amazon Chime SDK Media Insights Pipeline Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnMediaInsightsPipelineConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_chime.NewCfnMediaInsightsPipelineConfigurationPropsMixin(&CfnMediaInsightsPipelineConfigurationMixinProps{
//   	Elements: []interface{}{
//   		&MediaInsightsPipelineConfigurationElementProperty{
//   			AmazonTranscribeCallAnalyticsProcessorConfiguration: &AmazonTranscribeCallAnalyticsProcessorConfigurationProperty{
//   				CallAnalyticsStreamCategories: []*string{
//   					jsii.String("callAnalyticsStreamCategories"),
//   				},
//   				ContentIdentificationType: jsii.String("contentIdentificationType"),
//   				ContentRedactionType: jsii.String("contentRedactionType"),
//   				EnablePartialResultsStabilization: jsii.Boolean(false),
//   				FilterPartialResults: jsii.Boolean(false),
//   				LanguageCode: jsii.String("languageCode"),
//   				LanguageModelName: jsii.String("languageModelName"),
//   				PartialResultsStability: jsii.String("partialResultsStability"),
//   				PiiEntityTypes: jsii.String("piiEntityTypes"),
//   				PostCallAnalyticsSettings: &PostCallAnalyticsSettingsProperty{
//   					ContentRedactionOutput: jsii.String("contentRedactionOutput"),
//   					DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   					OutputEncryptionKmsKeyId: jsii.String("outputEncryptionKmsKeyId"),
//   					OutputLocation: jsii.String("outputLocation"),
//   				},
//   				VocabularyFilterMethod: jsii.String("vocabularyFilterMethod"),
//   				VocabularyFilterName: jsii.String("vocabularyFilterName"),
//   				VocabularyName: jsii.String("vocabularyName"),
//   			},
//   			AmazonTranscribeProcessorConfiguration: &AmazonTranscribeProcessorConfigurationProperty{
//   				ContentIdentificationType: jsii.String("contentIdentificationType"),
//   				ContentRedactionType: jsii.String("contentRedactionType"),
//   				EnablePartialResultsStabilization: jsii.Boolean(false),
//   				FilterPartialResults: jsii.Boolean(false),
//   				IdentifyLanguage: jsii.Boolean(false),
//   				IdentifyMultipleLanguages: jsii.Boolean(false),
//   				LanguageCode: jsii.String("languageCode"),
//   				LanguageModelName: jsii.String("languageModelName"),
//   				LanguageOptions: jsii.String("languageOptions"),
//   				PartialResultsStability: jsii.String("partialResultsStability"),
//   				PiiEntityTypes: jsii.String("piiEntityTypes"),
//   				PreferredLanguage: jsii.String("preferredLanguage"),
//   				ShowSpeakerLabel: jsii.Boolean(false),
//   				VocabularyFilterMethod: jsii.String("vocabularyFilterMethod"),
//   				VocabularyFilterName: jsii.String("vocabularyFilterName"),
//   				VocabularyFilterNames: jsii.String("vocabularyFilterNames"),
//   				VocabularyName: jsii.String("vocabularyName"),
//   				VocabularyNames: jsii.String("vocabularyNames"),
//   			},
//   			KinesisDataStreamSinkConfiguration: &KinesisDataStreamSinkConfigurationProperty{
//   				InsightsTarget: jsii.String("insightsTarget"),
//   			},
//   			S3RecordingSinkConfiguration: &S3RecordingSinkConfigurationProperty{
//   				Destination: jsii.String("destination"),
//   				RecordingFileFormat: jsii.String("recordingFileFormat"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	MediaInsightsPipelineConfigurationName: jsii.String("mediaInsightsPipelineConfigurationName"),
//   	RealTimeAlertConfiguration: &RealTimeAlertConfigurationProperty{
//   		Disabled: jsii.Boolean(false),
//   		Rules: []interface{}{
//   			&RealTimeAlertRuleProperty{
//   				IssueDetectionConfiguration: &IssueDetectionConfigurationProperty{
//   					RuleName: jsii.String("ruleName"),
//   				},
//   				KeywordMatchConfiguration: &KeywordMatchConfigurationProperty{
//   					Keywords: []*string{
//   						jsii.String("keywords"),
//   					},
//   					Negate: jsii.Boolean(false),
//   					RuleName: jsii.String("ruleName"),
//   				},
//   				SentimentConfiguration: &SentimentConfigurationProperty{
//   					RuleName: jsii.String("ruleName"),
//   					SentimentType: jsii.String("sentimentType"),
//   					TimePeriod: jsii.Number(123),
//   				},
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ResourceAccessRoleArn: jsii.String("resourceAccessRoleArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html
//
type CfnMediaInsightsPipelineConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnMediaInsightsPipelineConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnMediaInsightsPipelineConfigurationPropsMixin
type jsiiProxy_CfnMediaInsightsPipelineConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnMediaInsightsPipelineConfigurationPropsMixin) Props() *CfnMediaInsightsPipelineConfigurationMixinProps {
	var returns *CfnMediaInsightsPipelineConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnMediaInsightsPipelineConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Chime::MediaInsightsPipelineConfiguration`.
func NewCfnMediaInsightsPipelineConfigurationPropsMixin(props *CfnMediaInsightsPipelineConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnMediaInsightsPipelineConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnMediaInsightsPipelineConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnMediaInsightsPipelineConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnMediaInsightsPipelineConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Chime::MediaInsightsPipelineConfiguration`.
func NewCfnMediaInsightsPipelineConfigurationPropsMixin_Override(c CfnMediaInsightsPipelineConfigurationPropsMixin, props *CfnMediaInsightsPipelineConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnMediaInsightsPipelineConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnMediaInsightsPipelineConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnMediaInsightsPipelineConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnMediaInsightsPipelineConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnMediaInsightsPipelineConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_chime.CfnMediaInsightsPipelineConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnMediaInsightsPipelineConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnMediaInsightsPipelineConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

