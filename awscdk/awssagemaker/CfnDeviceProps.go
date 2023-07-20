package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDevice`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceProps := &CfnDeviceProps{
//   	DeviceFleetName: jsii.String("deviceFleetName"),
//
//   	// the properties below are optional
//   	Device: &DeviceProperty{
//   		DeviceName: jsii.String("deviceName"),
//
//   		// the properties below are optional
//   		Description: jsii.String("description"),
//   		IotThingName: jsii.String("iotThingName"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html
//
type CfnDeviceProps struct {
	// The name of the fleet the device belongs to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html#cfn-sagemaker-device-devicefleetname
	//
	DeviceFleetName *string `field:"required" json:"deviceFleetName" yaml:"deviceFleetName"`
	// Edge device you want to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html#cfn-sagemaker-device-device
	//
	Device interface{} `field:"optional" json:"device" yaml:"device"`
	// An array of key-value pairs that contain metadata to help you categorize and organize your devices.
	//
	// Each tag consists of a key and a value, both of which you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-device.html#cfn-sagemaker-device-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

