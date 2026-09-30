package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnMediaInsightsPipelineConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMediaInsightsPipelineConfigurationMixinProps := &CfnMediaInsightsPipelineConfigurationMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html
//
type CfnMediaInsightsPipelineConfigurationMixinProps struct {
	// The elements in the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-elements
	//
	Elements interface{} `field:"optional" json:"elements" yaml:"elements"`
	// The name of the media insights pipeline configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationname
	//
	MediaInsightsPipelineConfigurationName *string `field:"optional" json:"mediaInsightsPipelineConfigurationName" yaml:"mediaInsightsPipelineConfigurationName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertconfiguration
	//
	RealTimeAlertConfiguration interface{} `field:"optional" json:"realTimeAlertConfiguration" yaml:"realTimeAlertConfiguration"`
	// The ARN of the role used by the service to access Amazon Web Services resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-resourceaccessrolearn
	//
	ResourceAccessRoleArn *string `field:"optional" json:"resourceAccessRoleArn" yaml:"resourceAccessRoleArn"`
	// The tags associated with the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

