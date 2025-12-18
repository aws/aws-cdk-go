package awsbedrock


// Sets whether your project will process images or not.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageOverrideConfigurationProperty := &ImageOverrideConfigurationProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageoverrideconfiguration.html
//
type CfnDataAutomationProject_ImageOverrideConfigurationProperty struct {
	// Sets modality processing for image files.
	//
	// All modalities are enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageoverrideconfiguration.html#cfn-bedrock-dataautomationproject-imageoverrideconfiguration-modalityprocessing
	//
	ModalityProcessing interface{} `field:"optional" json:"modalityProcessing" yaml:"modalityProcessing"`
	// Configuration for sensitive data detection and redaction for image files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageoverrideconfiguration.html#cfn-bedrock-dataautomationproject-imageoverrideconfiguration-sensitivedataconfiguration
	//
	SensitiveDataConfiguration interface{} `field:"optional" json:"sensitiveDataConfiguration" yaml:"sensitiveDataConfiguration"`
}

