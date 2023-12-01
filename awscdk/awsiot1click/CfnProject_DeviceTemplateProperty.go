package awsiot1click


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var callbackOverrides interface{}
//
//   deviceTemplateProperty := &DeviceTemplateProperty{
//   	CallbackOverrides: callbackOverrides,
//   	DeviceType: jsii.String("deviceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-devicetemplate.html
//
type CfnProject_DeviceTemplateProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-devicetemplate.html#cfn-iot1click-project-devicetemplate-callbackoverrides
	//
	CallbackOverrides interface{} `field:"optional" json:"callbackOverrides" yaml:"callbackOverrides"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot1click-project-devicetemplate.html#cfn-iot1click-project-devicetemplate-devicetype
	//
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
}

