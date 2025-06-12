package awsbedrock


// Settings for generating data from images.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageStandardExtractionProperty := &ImageStandardExtractionProperty{
//   	BoundingBox: &ImageBoundingBoxProperty{
//   		State: jsii.String("state"),
//   	},
//   	Category: &ImageExtractionCategoryProperty{
//   		State: jsii.String("state"),
//
//   		// the properties below are optional
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardextraction.html
//
type CfnDataAutomationProject_ImageStandardExtractionProperty struct {
	// Settings for generating bounding boxes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardextraction.html#cfn-bedrock-dataautomationproject-imagestandardextraction-boundingbox
	//
	BoundingBox interface{} `field:"required" json:"boundingBox" yaml:"boundingBox"`
	// Settings for generating categorical data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imagestandardextraction.html#cfn-bedrock-dataautomationproject-imagestandardextraction-category
	//
	Category interface{} `field:"required" json:"category" yaml:"category"`
}

