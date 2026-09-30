package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonTranscribeCallAnalyticsProcessorConfigurationProperty := &AmazonTranscribeCallAnalyticsProcessorConfigurationProperty{
//   	LanguageCode: jsii.String("languageCode"),
//
//   	// the properties below are optional
//   	CallAnalyticsStreamCategories: []*string{
//   		jsii.String("callAnalyticsStreamCategories"),
//   	},
//   	ContentIdentificationType: jsii.String("contentIdentificationType"),
//   	ContentRedactionType: jsii.String("contentRedactionType"),
//   	EnablePartialResultsStabilization: jsii.Boolean(false),
//   	FilterPartialResults: jsii.Boolean(false),
//   	LanguageModelName: jsii.String("languageModelName"),
//   	PartialResultsStability: jsii.String("partialResultsStability"),
//   	PiiEntityTypes: jsii.String("piiEntityTypes"),
//   	PostCallAnalyticsSettings: &PostCallAnalyticsSettingsProperty{
//   		DataAccessRoleArn: jsii.String("dataAccessRoleArn"),
//   		OutputLocation: jsii.String("outputLocation"),
//
//   		// the properties below are optional
//   		ContentRedactionOutput: jsii.String("contentRedactionOutput"),
//   		OutputEncryptionKmsKeyId: jsii.String("outputEncryptionKmsKeyId"),
//   	},
//   	VocabularyFilterMethod: jsii.String("vocabularyFilterMethod"),
//   	VocabularyFilterName: jsii.String("vocabularyFilterName"),
//   	VocabularyName: jsii.String("vocabularyName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html
//
type CfnMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty struct {
	// The language code in the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-languagecode
	//
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// The categories to send to the insights target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-callanalyticsstreamcategories
	//
	CallAnalyticsStreamCategories *[]*string `field:"optional" json:"callAnalyticsStreamCategories" yaml:"callAnalyticsStreamCategories"`
	// Labels all PII identified in the transcript.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-contentidentificationtype
	//
	ContentIdentificationType *string `field:"optional" json:"contentIdentificationType" yaml:"contentIdentificationType"`
	// Redacts all PII identified in the transcript.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-contentredactiontype
	//
	ContentRedactionType *string `field:"optional" json:"contentRedactionType" yaml:"contentRedactionType"`
	// Enables partial result stabilization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-enablepartialresultsstabilization
	//
	EnablePartialResultsStabilization interface{} `field:"optional" json:"enablePartialResultsStabilization" yaml:"enablePartialResultsStabilization"`
	// If true, partial results are filtered out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-filterpartialresults
	//
	FilterPartialResults interface{} `field:"optional" json:"filterPartialResults" yaml:"filterPartialResults"`
	// The name of the custom language model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-languagemodelname
	//
	LanguageModelName *string `field:"optional" json:"languageModelName" yaml:"languageModelName"`
	// The level of stability for partial results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-partialresultsstability
	//
	PartialResultsStability *string `field:"optional" json:"partialResultsStability" yaml:"partialResultsStability"`
	// The types of PII to redact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-piientitytypes
	//
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-postcallanalyticssettings
	//
	PostCallAnalyticsSettings interface{} `field:"optional" json:"postCallAnalyticsSettings" yaml:"postCallAnalyticsSettings"`
	// The vocabulary filtering method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-vocabularyfiltermethod
	//
	VocabularyFilterMethod *string `field:"optional" json:"vocabularyFilterMethod" yaml:"vocabularyFilterMethod"`
	// The name of the custom vocabulary filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-vocabularyfiltername
	//
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// The name of the custom vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribecallanalyticsprocessorconfiguration-vocabularyname
	//
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
}

