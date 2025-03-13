package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageStandardOutputConfigurationProperty := &ImageStandardOutputConfigurationProperty{
//   	Extraction: &ImageStandardExtractionProperty{
//   		BoundingBox: &ImageBoundingBoxProperty{
//   			State: jsii.String("state"),
//   		},
//   		Category: &ImageExtractionCategoryProperty{
//   			State: jsii.String("state"),
//
//   			// the properties below are optional
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   	GenerativeField: &ImageStandardGenerativeFieldProperty{
//   		State: jsii.String("state"),
//
//   		// the properties below are optional
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardoutputconfiguration.html
//
type CfnDataAutomationProject_ImageStandardOutputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-imagestandardoutputconfiguration-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-imagestandardoutputconfiguration-generativefield
	//
	GenerativeField interface{} `field:"optional" json:"generativeField" yaml:"generativeField"`
}

