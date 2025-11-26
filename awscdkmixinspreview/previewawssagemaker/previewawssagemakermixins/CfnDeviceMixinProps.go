package previewawssagemakermixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDevicePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeviceMixinProps := &CfnDeviceMixinProps{
//   	Device: &DeviceProperty{
//   		Description: jsii.String("description"),
//   		DeviceName: jsii.String("deviceName"),
//   		IotThingName: jsii.String("iotThingName"),
//   	},
//   	DeviceFleetName: jsii.String("deviceFleetName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html
//
type CfnDeviceMixinProps struct {
	// Edge device you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html#cfn-sagemaker-device-device
	//
	Device interface{} `field:"optional" json:"device" yaml:"device"`
	// The name of the fleet the device belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html#cfn-sagemaker-device-devicefleetname
	//
	DeviceFleetName *string `field:"optional" json:"deviceFleetName" yaml:"deviceFleetName"`
	// An array of key-value pairs that contain metadata to help you categorize and organize your devices.
	//
	// Each tag consists of a key and a value, both of which you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html#cfn-sagemaker-device-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

