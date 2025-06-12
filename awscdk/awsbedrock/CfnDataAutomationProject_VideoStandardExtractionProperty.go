package awsbedrock


// Settings for generating data from video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoStandardExtractionProperty := &VideoStandardExtractionProperty{
//   	BoundingBox: &VideoBoundingBoxProperty{
//   		State: jsii.String("state"),
//   	},
//   	Category: &VideoExtractionCategoryProperty{
//   		State: jsii.String("state"),
//
//   		// the properties below are optional
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardextraction.html
//
type CfnDataAutomationProject_VideoStandardExtractionProperty struct {
	// Settings for generating bounding boxes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardextraction.html#cfn-bedrock-dataautomationproject-videostandardextraction-boundingbox
	//
	BoundingBox interface{} `field:"required" json:"boundingBox" yaml:"boundingBox"`
	// Settings for generating categorical data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videostandardextraction.html#cfn-bedrock-dataautomationproject-videostandardextraction-category
	//
	Category interface{} `field:"required" json:"category" yaml:"category"`
}

