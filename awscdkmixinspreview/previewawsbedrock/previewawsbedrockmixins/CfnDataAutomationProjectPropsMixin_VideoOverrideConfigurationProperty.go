package previewawsbedrockmixins


// Sets whether your project will process videos or not.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   videoOverrideConfigurationProperty := &VideoOverrideConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videooverrideconfiguration.html
//
type CfnDataAutomationProjectPropsMixin_VideoOverrideConfigurationProperty struct {
	// Sets modality processing for video files.
	//
	// All modalities are enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videooverrideconfiguration.html#cfn-bedrock-dataautomationproject-videooverrideconfiguration-modalityprocessing
	//
	ModalityProcessing interface{} `field:"optional" json:"modalityProcessing" yaml:"modalityProcessing"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videooverrideconfiguration.html#cfn-bedrock-dataautomationproject-videooverrideconfiguration-sensitivedataconfiguration
	//
	SensitiveDataConfiguration interface{} `field:"optional" json:"sensitiveDataConfiguration" yaml:"sensitiveDataConfiguration"`
}

