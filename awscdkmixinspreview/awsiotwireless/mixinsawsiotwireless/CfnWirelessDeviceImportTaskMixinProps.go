package mixinsawsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWirelessDeviceImportTaskPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnWirelessDeviceImportTaskMixinProps := &CfnWirelessDeviceImportTaskMixinProps{
//   	DestinationName: jsii.String("destinationName"),
//   	Sidewalk: &SidewalkProperty{
//   		DeviceCreationFile: jsii.String("deviceCreationFile"),
//   		DeviceCreationFileList: []*string{
//   			jsii.String("deviceCreationFileList"),
//   		},
//   		Role: jsii.String("role"),
//   		SidewalkManufacturingSn: jsii.String("sidewalkManufacturingSn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdeviceimporttask.html
//
type CfnWirelessDeviceImportTaskMixinProps struct {
	// The name of the destination that describes the IoT rule to route messages from the Sidewalk devices in the import task to other applications.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdeviceimporttask.html#cfn-iotwireless-wirelessdeviceimporttask-destinationname
	//
	DestinationName *string `field:"optional" json:"destinationName" yaml:"destinationName"`
	// The Sidewalk-related information of the wireless device import task.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdeviceimporttask.html#cfn-iotwireless-wirelessdeviceimporttask-sidewalk
	//
	Sidewalk interface{} `field:"optional" json:"sidewalk" yaml:"sidewalk"`
	// Adds to or modifies the tags of the given resource.
	//
	// Tags are metadata that you can use to manage a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotwireless-wirelessdeviceimporttask.html#cfn-iotwireless-wirelessdeviceimporttask-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

