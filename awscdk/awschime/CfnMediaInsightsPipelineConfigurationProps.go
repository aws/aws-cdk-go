package awschime

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnMediaInsightsPipelineConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMediaInsightsPipelineConfigurationProps := &CfnMediaInsightsPipelineConfigurationProps{
//   	Elements: []interface{}{
//   		&MediaInsightsPipelineConfigurationElementProperty{
//   			Type: jsii.String("type"),
//
//   			// the properties below are optional
//   			AmazonTranscribeCallAnalyticsProcessorConfiguration: &AmazonTranscribeCallAnalyticsProcessorConfigurationProperty{
//   				LanguageCode: jsii.String("languageCode"),
//
//   				// the properties below are optional
//   				CallAnalyticsStreamCategories: []*string{
//   					jsii.String("callAnalyticsStreamCategories"),
//   				},
//   				ContentIdentificationType: jsii.String("contentIdentificationType"),
//   				ContentRedactionType: jsii.String("contentRedactionType"),
//   				EnablePartialResultsStabilization: jsii.Boolean(false),
//   				FilterPartialResults: jsii.Boolean(false),
//   				LanguageModelName: jsii.String("languageModelName"),
//   				PartialResultsStability: jsii.String("partialResultsStability"),
//   				PiiEntityTypes: jsii.String("piiEntityTypes"),
//   				PostCallAnalyticsSettings: &PostCallAnalyticsSettingsProperty{
//   					DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   					OutputLocation: jsii.String("outputLocation"),
//
//   					// the properties below are optional
//   					ContentRedactionOutput: jsii.String("contentRedactionOutput"),
//   					OutputEncryptionKmsKeyId: jsii.String("outputEncryptionKmsKeyId"),
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
//   		},
//   	},
//   	MediaInsightsPipelineConfigurationName: jsii.String("mediaInsightsPipelineConfigurationName"),
//   	ResourceAccessRoleArn: jsii.String("resourceAccessRoleArn"),
//
//   	// the properties below are optional
//   	RealTimeAlertConfiguration: &RealTimeAlertConfigurationProperty{
//   		Disabled: jsii.Boolean(false),
//   		Rules: []interface{}{
//   			&RealTimeAlertRuleProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				IssueDetectionConfiguration: &IssueDetectionConfigurationProperty{
//   					RuleName: jsii.String("ruleName"),
//   				},
//   				KeywordMatchConfiguration: &KeywordMatchConfigurationProperty{
//   					Keywords: []*string{
//   						jsii.String("keywords"),
//   					},
//   					RuleName: jsii.String("ruleName"),
//
//   					// the properties below are optional
//   					Negate: jsii.Boolean(false),
//   				},
//   				SentimentConfiguration: &SentimentConfigurationProperty{
//   					RuleName: jsii.String("ruleName"),
//   					SentimentType: jsii.String("sentimentType"),
//   					TimePeriod: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
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
type CfnMediaInsightsPipelineConfigurationProps struct {
	// The elements in the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-elements
	//
	Elements interface{} `field:"required" json:"elements" yaml:"elements"`
	// The name of the media insights pipeline configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationname
	//
	MediaInsightsPipelineConfigurationName *string `field:"required" json:"mediaInsightsPipelineConfigurationName" yaml:"mediaInsightsPipelineConfigurationName"`
	// The ARN of the role used by the service to access Amazon Web Services resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-resourceaccessrolearn
	//
	ResourceAccessRoleArn *string `field:"required" json:"resourceAccessRoleArn" yaml:"resourceAccessRoleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-realtimealertconfiguration
	//
	RealTimeAlertConfiguration interface{} `field:"optional" json:"realTimeAlertConfiguration" yaml:"realTimeAlertConfiguration"`
	// The tags associated with the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-chime-mediainsightspipelineconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

