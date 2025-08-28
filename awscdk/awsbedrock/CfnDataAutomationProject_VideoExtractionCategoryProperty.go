package awsbedrock


// Settings for generating categorical data from video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoExtractionCategoryProperty := &VideoExtractionCategoryProperty{
//   	State: jsii.String("state"),
//
//   	// the properties below are optional
//   	Types: []*string{
//   		jsii.String("types"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videoextractioncategory.html
//
type CfnDataAutomationProject_VideoExtractionCategoryProperty struct {
	// Whether generating categorical data from video is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videoextractioncategory.html#cfn-bedrock-dataautomationproject-videoextractioncategory-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
	// The types of data to generate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videoextractioncategory.html#cfn-bedrock-dataautomationproject-videoextractioncategory-types
	//
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

