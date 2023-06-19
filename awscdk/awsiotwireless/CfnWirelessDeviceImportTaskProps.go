package awsiotwireless

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnWirelessDeviceImportTask`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWirelessDeviceImportTaskProps := &CfnWirelessDeviceImportTaskProps{
//   	DestinationName: jsii.String("destinationName"),
//   	Sidewalk: &SidewalkProperty{
//   		DeviceCreationFile: jsii.String("deviceCreationFile"),
//   		DeviceCreationFileList: []*string{
//   			jsii.String("deviceCreationFileList"),
//   		},
//   		Role: jsii.String("role"),
//   		SidewalkManufacturingSn: jsii.String("sidewalkManufacturingSn"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnWirelessDeviceImportTaskProps struct {
	// The name of the destination that describes the IoT rule to route messages from the Sidewalk devices in the import task to other applications.
	DestinationName *string `field:"required" json:"destinationName" yaml:"destinationName"`
	// The Sidewalk-related information of the wireless device import task.
	Sidewalk interface{} `field:"required" json:"sidewalk" yaml:"sidewalk"`
	// Adds to or modifies the tags of the given resource.
	//
	// Tags are metadata that you can use to manage a resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

