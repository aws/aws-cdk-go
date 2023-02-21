package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDeviceFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceFleetProps := &cfnDeviceFleetProps{
//   	deviceFleetName: jsii.String("deviceFleetName"),
//   	outputConfig: &edgeOutputConfigProperty{
//   		s3OutputLocation: jsii.String("s3OutputLocation"),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDeviceFleetProps struct {
	// Name of the device fleet.
	DeviceFleetName *string `field:"required" json:"deviceFleetName" yaml:"deviceFleetName"`
	// The output configuration for storing sample data collected by the fleet.
	OutputConfig interface{} `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A description of the fleet.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An array of key-value pairs that contain metadata to help you categorize and organize your device fleets.
	//
	// Each tag consists of a key and a value, both of which you define.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

