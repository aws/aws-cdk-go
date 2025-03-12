package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   audioStandardExtractionProperty := &AudioStandardExtractionProperty{
//   	Category: &AudioExtractionCategoryProperty{
//   		State: jsii.String("state"),
//
//   		// the properties below are optional
//   		Types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardextraction.html
//
type CfnDataAutomationProject_AudioStandardExtractionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-audiostandardextraction.html#cfn-bedrock-dataautomationproject-audiostandardextraction-category
	//
	Category interface{} `field:"required" json:"category" yaml:"category"`
}

