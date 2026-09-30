package awschime


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amazonTranscribeProcessorConfigurationProperty := &AmazonTranscribeProcessorConfigurationProperty{
//   	ContentIdentificationType: jsii.String("contentIdentificationType"),
//   	ContentRedactionType: jsii.String("contentRedactionType"),
//   	EnablePartialResultsStabilization: jsii.Boolean(false),
//   	FilterPartialResults: jsii.Boolean(false),
//   	IdentifyLanguage: jsii.Boolean(false),
//   	IdentifyMultipleLanguages: jsii.Boolean(false),
//   	LanguageCode: jsii.String("languageCode"),
//   	LanguageModelName: jsii.String("languageModelName"),
//   	LanguageOptions: jsii.String("languageOptions"),
//   	PartialResultsStability: jsii.String("partialResultsStability"),
//   	PiiEntityTypes: jsii.String("piiEntityTypes"),
//   	PreferredLanguage: jsii.String("preferredLanguage"),
//   	ShowSpeakerLabel: jsii.Boolean(false),
//   	VocabularyFilterMethod: jsii.String("vocabularyFilterMethod"),
//   	VocabularyFilterName: jsii.String("vocabularyFilterName"),
//   	VocabularyFilterNames: jsii.String("vocabularyFilterNames"),
//   	VocabularyName: jsii.String("vocabularyName"),
//   	VocabularyNames: jsii.String("vocabularyNames"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html
//
type CfnMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty struct {
	// Labels all PII identified in the transcript.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-contentidentificationtype
	//
	ContentIdentificationType *string `field:"optional" json:"contentIdentificationType" yaml:"contentIdentificationType"`
	// Redacts all PII identified in the transcript.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-contentredactiontype
	//
	ContentRedactionType *string `field:"optional" json:"contentRedactionType" yaml:"contentRedactionType"`
	// Enables partial result stabilization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-enablepartialresultsstabilization
	//
	EnablePartialResultsStabilization interface{} `field:"optional" json:"enablePartialResultsStabilization" yaml:"enablePartialResultsStabilization"`
	// If true, partial results are filtered out.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-filterpartialresults
	//
	FilterPartialResults interface{} `field:"optional" json:"filterPartialResults" yaml:"filterPartialResults"`
	// Turns language identification on or off.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-identifylanguage
	//
	IdentifyLanguage interface{} `field:"optional" json:"identifyLanguage" yaml:"identifyLanguage"`
	// Turns multiple language identification on or off.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-identifymultiplelanguages
	//
	IdentifyMultipleLanguages interface{} `field:"optional" json:"identifyMultipleLanguages" yaml:"identifyMultipleLanguages"`
	// The language code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// The name of the custom language model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-languagemodelname
	//
	LanguageModelName *string `field:"optional" json:"languageModelName" yaml:"languageModelName"`
	// The language options for transcription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-languageoptions
	//
	LanguageOptions *string `field:"optional" json:"languageOptions" yaml:"languageOptions"`
	// The level of stability for partial results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-partialresultsstability
	//
	PartialResultsStability *string `field:"optional" json:"partialResultsStability" yaml:"partialResultsStability"`
	// The types of PII to redact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-piientitytypes
	//
	PiiEntityTypes *string `field:"optional" json:"piiEntityTypes" yaml:"piiEntityTypes"`
	// The preferred language for transcription.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-preferredlanguage
	//
	PreferredLanguage *string `field:"optional" json:"preferredLanguage" yaml:"preferredLanguage"`
	// Enables speaker partitioning.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-showspeakerlabel
	//
	ShowSpeakerLabel interface{} `field:"optional" json:"showSpeakerLabel" yaml:"showSpeakerLabel"`
	// The vocabulary filtering method.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-vocabularyfiltermethod
	//
	VocabularyFilterMethod *string `field:"optional" json:"vocabularyFilterMethod" yaml:"vocabularyFilterMethod"`
	// The name of the custom vocabulary filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-vocabularyfiltername
	//
	VocabularyFilterName *string `field:"optional" json:"vocabularyFilterName" yaml:"vocabularyFilterName"`
	// The names of the custom vocabulary filters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-vocabularyfilternames
	//
	VocabularyFilterNames *string `field:"optional" json:"vocabularyFilterNames" yaml:"vocabularyFilterNames"`
	// The name of the custom vocabulary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-vocabularyname
	//
	VocabularyName *string `field:"optional" json:"vocabularyName" yaml:"vocabularyName"`
	// The names of the custom vocabularies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration.html#cfn-chime-mediainsightspipelineconfiguration-amazontranscribeprocessorconfiguration-vocabularynames
	//
	VocabularyNames *string `field:"optional" json:"vocabularyNames" yaml:"vocabularyNames"`
}

