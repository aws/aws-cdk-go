package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inferenceAcceleratorProperty := &InferenceAcceleratorProperty{
//   	DeviceName: jsii.String("deviceName"),
//   	DeviceType: jsii.String("deviceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-inferenceaccelerator.html
//
type CfnTaskDefinitionPropsMixin_InferenceAcceleratorProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-inferenceaccelerator.html#cfn-ecs-taskdefinition-inferenceaccelerator-devicename
	//
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-inferenceaccelerator.html#cfn-ecs-taskdefinition-inferenceaccelerator-devicetype
	//
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

