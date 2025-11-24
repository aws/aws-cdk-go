package mixinsawsbedrock


// Bounding box settings for a project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageBoundingBoxProperty := &ImageBoundingBoxProperty{
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageboundingbox.html
//
type CfnDataAutomationProjectPropsMixin_ImageBoundingBoxProperty struct {
	// Bounding box settings for a project.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-dataautomationproject-imageboundingbox.html#cfn-bedrock-dataautomationproject-imageboundingbox-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}

