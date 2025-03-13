package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentStandardOutputConfigurationProperty := &DocumentStandardOutputConfigurationProperty{
//   	Extraction: &DocumentStandardExtractionProperty{
//   		BoundingBox: &DocumentBoundingBoxProperty{
//   			State: jsii.String("state"),
//   		},
//   		Granularity: &DocumentExtractionGranularityProperty{
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   	GenerativeField: &DocumentStandardGenerativeFieldProperty{
//   		State: jsii.String("state"),
//   	},
//   	OutputFormat: &DocumentOutputFormatProperty{
//   		AdditionalFileFormat: &DocumentOutputAdditionalFileFormatProperty{
//   			State: jsii.String("state"),
//   		},
//   		TextFormat: &DocumentOutputTextFormatProperty{
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardoutputconfiguration.html
//
type CfnDataAutomationProject_DocumentStandardOutputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-documentstandardoutputconfiguration-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-documentstandardoutputconfiguration-generativefield
	//
	GenerativeField interface{} `field:"optional" json:"generativeField" yaml:"generativeField"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-documentstandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-documentstandardoutputconfiguration-outputformat
	//
	OutputFormat interface{} `field:"optional" json:"outputFormat" yaml:"outputFormat"`
}

