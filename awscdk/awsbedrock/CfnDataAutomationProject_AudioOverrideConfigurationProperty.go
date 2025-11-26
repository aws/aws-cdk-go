package awsbedrock


// Sets whether your project will process audio or not.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioOverrideConfigurationProperty := &AudioOverrideConfigurationProperty{
//   	LanguageConfiguration: &AudioLanguageConfigurationProperty{
//   		GenerativeOutputLanguage: jsii.String("generativeOutputLanguage"),
//   		IdentifyMultipleLanguages: jsii.Boolean(false),
//   		InputLanguages: []*string{
//   			jsii.String("inputLanguages"),
//   		},
//   	},
//   	ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   		State: jsii.String("state"),
//   	},
//   	SensitiveDataConfiguration: &SensitiveDataConfigurationProperty{
//   		DetectionMode: jsii.String("detectionMode"),
//   		DetectionScope: []*string{
//   			jsii.String("detectionScope"),
//   		},
//   		PiiEntitiesConfiguration: &PIIEntitiesConfigurationProperty{
//   			PiiEntityTypes: []*string{
//   				jsii.String("piiEntityTypes"),
//   			},
//   			RedactionMaskMode: jsii.String("redactionMaskMode"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiooverrideconfiguration.html
//
type CfnDataAutomationProject_AudioOverrideConfigurationProperty struct {
	// The output and input language configuration for your audio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiooverrideconfiguration.html#cfn-bedrock-dataautomationproject-audiooverrideconfiguration-languageconfiguration
	//
	LanguageConfiguration interface{} `field:"optional" json:"languageConfiguration" yaml:"languageConfiguration"`
	// Sets modality processing for audio files.
	//
	// All modalities are enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiooverrideconfiguration.html#cfn-bedrock-dataautomationproject-audiooverrideconfiguration-modalityprocessing
	//
	ModalityProcessing interface{} `field:"optional" json:"modalityProcessing" yaml:"modalityProcessing"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiooverrideconfiguration.html#cfn-bedrock-dataautomationproject-audiooverrideconfiguration-sensitivedataconfiguration
	//
	SensitiveDataConfiguration interface{} `field:"optional" json:"sensitiveDataConfiguration" yaml:"sensitiveDataConfiguration"`
}

