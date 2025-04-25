package awsbedrock


// Bounding box settings for video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   videoBoundingBoxProperty := &VideoBoundingBoxProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videoboundingbox.html
//
type CfnDataAutomationProject_VideoBoundingBoxProperty struct {
	// Whether bounding boxes are enabled for video.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-videoboundingbox.html#cfn-bedrock-dataautomationproject-videoboundingbox-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

