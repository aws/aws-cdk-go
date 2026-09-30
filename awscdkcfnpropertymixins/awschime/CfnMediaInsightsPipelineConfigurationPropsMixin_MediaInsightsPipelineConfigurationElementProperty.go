package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   mediaInsightsPipelineConfigurationElementProperty := &MediaInsightsPipelineConfigurationElementProperty{
//   	AmazonTranscribeCallAnalyticsProcessorConfiguration: &AmazonTranscribeCallAnalyticsProcessorConfigurationProperty{
//   		CallAnalyticsStreamCategories: []*string{
//   			jsii.String("callAnalyticsStreamCategories"),
//   		},
//   		ContentIdentificationType: jsii.String("contentIdentificationType"),
//   		ContentRedactionType: jsii.String("contentRedactionType"),
//   		EnablePartialResultsStabilization: jsii.Boolean(false),
//   		FilterPartialResults: jsii.Boolean(false),
//   		LanguageCode: jsii.String("languageCode"),
//   		LanguageModelName: jsii.String("languageModelName"),
//   		PartialResultsStability: jsii.String("partialResultsStability"),
//   		PiiEntityTypes: jsii.String("piiEntityTypes"),
//   		PostCallAnalyticsSettings: &PostCallAnalyticsSettingsProperty{
//   			ContentRedactionOutput: jsii.String("contentRedactionOutput"),
//   			DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   			OutputEncryptionKmsKeyId: jsii.String("outputEncryptionKmsKeyId"),
//   			OutputLocation: jsii.String("outputLocation"),
//   		},
//   		VocabularyFilterMethod: jsii.String("vocabularyFilterMethod"),
//   		VocabularyFilterName: jsii.String("vocabularyFilterName"),
//   		VocabularyName: jsii.String("vocabularyName"),
//   	},
//   	AmazonTranscribeProcessorConfiguration: &AmazonTranscribeProcessorConfigurationProperty{
//   		ContentIdentificationType: jsii.String("contentIdentificationType"),
//   		ContentRedactionType: jsii.String("contentRedactionType"),
//   		EnablePartialResultsStabilization: jsii.Boolean(false),
//   		FilterPartialResults: jsii.Boolean(false),
//   		IdentifyLanguage: jsii.Boolean(false),
//   		IdentifyMultipleLanguages: jsii.Boolean(false),
//   		LanguageCode: jsii.String("languageCode"),
//   		LanguageModelName: jsii.String("languageModelName"),
//   		LanguageOptions: jsii.String("languageOptions"),
//   		PartialResultsStability: jsii.String("partialResultsStability"),
//   		PiiEntityTypes: jsii.String("piiEntityTypes"),
//   		PreferredLanguage: jsii.String("preferredLanguage"),
//   		ShowSpeakerLabel: jsii.Boolean(false),
//   		VocabularyFilterMethod: jsii.String("vocabularyFilterMethod"),
//   		VocabularyFilterName: jsii.String("vocabularyFilterName"),
//   		VocabularyFilterNames: jsii.String("vocabularyFilterNames"),
//   		VocabularyName: jsii.String("vocabularyName"),
//   		VocabularyNames: jsii.String("vocabularyNames"),
//   	},
//   	KinesisDataStreamSinkConfiguration: &KinesisDataStreamSinkConfigurationProperty{
//   		InsightsTarget: jsii.String("insightsTarget"),
//   	},
//   	S3RecordingSinkConfiguration: &S3RecordingSinkConfigurationProperty{
//   		Destination: jsii.String("destination"),
//   		RecordingFileFormat: jsii.String("recordingFileFormat"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement.html
//
type CfnMediaInsightsPipelineConfigurationPropsMixin_MediaInsightsPipelineConfigurationElementProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement.html#cfn-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement-amazontranscribecallanalyticsprocessorconfiguration
	//
	AmazonTranscribeCallAnalyticsProcessorConfiguration interface{} `field:"optional" json:"amazonTranscribeCallAnalyticsProcessorConfiguration" yaml:"amazonTranscribeCallAnalyticsProcessorConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement.html#cfn-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement-amazontranscribeprocessorconfiguration
	//
	AmazonTranscribeProcessorConfiguration interface{} `field:"optional" json:"amazonTranscribeProcessorConfiguration" yaml:"amazonTranscribeProcessorConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement.html#cfn-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement-kinesisdatastreamsinkconfiguration
	//
	KinesisDataStreamSinkConfiguration interface{} `field:"optional" json:"kinesisDataStreamSinkConfiguration" yaml:"kinesisDataStreamSinkConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement.html#cfn-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement-s3recordingsinkconfiguration
	//
	S3RecordingSinkConfiguration interface{} `field:"optional" json:"s3RecordingSinkConfiguration" yaml:"s3RecordingSinkConfiguration"`
	// The element type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement.html#cfn-chime-mediainsightspipelineconfiguration-mediainsightspipelineconfigurationelement-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

