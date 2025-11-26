package awsbedrock


// Additional settings for a project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentOverrideConfigurationProperty := &DocumentOverrideConfigurationProperty{
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
//   	Splitter: &SplitterConfigurationProperty{
//   		State: jsii.String("state"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoverrideconfiguration.html
//
type CfnDataAutomationProject_DocumentOverrideConfigurationProperty struct {
	// Sets modality processing for document files.
	//
	// All modalities are enabled by default.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoverrideconfiguration.html#cfn-bedrock-dataautomationproject-documentoverrideconfiguration-modalityprocessing
	//
	ModalityProcessing interface{} `field:"optional" json:"modalityProcessing" yaml:"modalityProcessing"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoverrideconfiguration.html#cfn-bedrock-dataautomationproject-documentoverrideconfiguration-sensitivedataconfiguration
	//
	SensitiveDataConfiguration interface{} `field:"optional" json:"sensitiveDataConfiguration" yaml:"sensitiveDataConfiguration"`
	// Whether document splitter is enabled for a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentoverrideconfiguration.html#cfn-bedrock-dataautomationproject-documentoverrideconfiguration-splitter
	//
	Splitter interface{} `field:"optional" json:"splitter" yaml:"splitter"`
}

