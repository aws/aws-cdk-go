package awsbedrock


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   imageBoundingBoxProperty := &ImageBoundingBoxProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageboundingbox.html
//
type CfnDataAutomationProject_ImageBoundingBoxProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageboundingbox.html#cfn-bedrock-dataautomationproject-imageboundingbox-state
	//
	State *string `field:"required" json:"state" yaml:"state"`
}

