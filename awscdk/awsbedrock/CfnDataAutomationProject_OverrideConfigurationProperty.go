package awsbedrock


// Additional settings for a project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   overrideConfigurationProperty := &OverrideConfigurationProperty{
//   	Audio: &AudioOverrideConfigurationProperty{
//   		ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   			State: jsii.String("state"),
//   		},
//   	},
//   	Document: &DocumentOverrideConfigurationProperty{
//   		ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   			State: jsii.String("state"),
//   		},
//   		Splitter: &SplitterConfigurationProperty{
//   			State: jsii.String("state"),
//   		},
//   	},
//   	Image: &ImageOverrideConfigurationProperty{
//   		ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   			State: jsii.String("state"),
//   		},
//   	},
//   	ModalityRouting: &ModalityRoutingConfigurationProperty{
//   		Jpeg: jsii.String("jpeg"),
//   		Mov: jsii.String("mov"),
//   		Mp4: jsii.String("mp4"),
//   		Png: jsii.String("png"),
//   	},
//   	Video: &VideoOverrideConfigurationProperty{
//   		ModalityProcessing: &ModalityProcessingConfigurationProperty{
//   			State: jsii.String("state"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html
//
type CfnDataAutomationProject_OverrideConfigurationProperty struct {
	// This element declares whether your project will process audio files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html#cfn-bedrock-dataautomationproject-overrideconfiguration-audio
	//
	Audio interface{} `field:"optional" json:"audio" yaml:"audio"`
	// Additional settings for a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html#cfn-bedrock-dataautomationproject-overrideconfiguration-document
	//
	Document interface{} `field:"optional" json:"document" yaml:"document"`
	// This element declares whether your project will process image files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html#cfn-bedrock-dataautomationproject-overrideconfiguration-image
	//
	Image interface{} `field:"optional" json:"image" yaml:"image"`
	// Lets you set which modalities certain file types are processed as.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html#cfn-bedrock-dataautomationproject-overrideconfiguration-modalityrouting
	//
	ModalityRouting interface{} `field:"optional" json:"modalityRouting" yaml:"modalityRouting"`
	// This element declares whether your project will process video files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-overrideconfiguration.html#cfn-bedrock-dataautomationproject-overrideconfiguration-video
	//
	Video interface{} `field:"optional" json:"video" yaml:"video"`
}

