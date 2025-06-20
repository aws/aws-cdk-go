package awsbedrock


// Output settings for processing video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoStandardOutputConfigurationProperty := &VideoStandardOutputConfigurationProperty{
//   	Extraction: &VideoStandardExtractionProperty{
//   		BoundingBox: &VideoBoundingBoxProperty{
//   			State: jsii.String("state"),
//   		},
//   		Category: &VideoExtractionCategoryProperty{
//   			State: jsii.String("state"),
//
//   			// the properties below are optional
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   	GenerativeField: &VideoStandardGenerativeFieldProperty{
//   		State: jsii.String("state"),
//
//   		// the properties below are optional
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardoutputconfiguration.html
//
type CfnDataAutomationProject_VideoStandardOutputConfigurationProperty struct {
	// Settings for populating data fields that describe the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-videostandardoutputconfiguration-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// Whether to generate descriptions of the video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardoutputconfiguration.html#cfn-bedrock-dataautomationproject-videostandardoutputconfiguration-generativefield
	//
	GenerativeField interface{} `field:"optional" json:"generativeField" yaml:"generativeField"`
}

