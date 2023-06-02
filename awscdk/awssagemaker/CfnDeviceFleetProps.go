package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDeviceFleet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceFleetProps := &CfnDeviceFleetProps{
//   	DeviceFleetName: jsii.String("deviceFleetName"),
//   	OutputConfig: &EdgeOutputConfigProperty{
//   		S3OutputLocation: jsii.String("s3OutputLocation"),
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

