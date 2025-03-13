package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioStandardOutputConfigurationProperty := &AudioStandardOutputConfigurationProperty{
//   	Extraction: &AudioStandardExtractionProperty{
//   		Category: &AudioExtractionCategoryProperty{
//   			State: jsii.String("state"),
//
//   			// the properties below are optional
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   	GenerativeField: &AudioStandardGenerativeFieldProperty{
//   		State: jsii.String("state"),
//
//   		// the properties below are optional
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardoutputconfiguration.html
//
type CfnDataAutomationProject_AudioStandardOutputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-audiostandardoutputconfiguration-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-audiostandardoutputconfiguration-generativefield
	//
	GenerativeField interface{} `field:"optional" json:"generativeField" yaml:"generativeField"`
}

