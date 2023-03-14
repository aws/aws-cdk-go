package awsgreengrass


// Properties for defining a `CfnDeviceDefinitionVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeviceDefinitionVersionProps := &cfnDeviceDefinitionVersionProps{
//   	deviceDefinitionId: jsii.String("deviceDefinitionId"),
//   	devices: []interface{}{
//   		&deviceProperty{
//   			certificateArn: jsii.String("certificateArn"),
//   			id: jsii.String("id"),
//   			thingArn: jsii.String("thingArn"),
//
//   			// the properties below are optional
//   			syncShadow: jsii.Boolean(false),
//   		},
//   	},
//   }
//
type CfnDeviceDefinitionVersionProps struct {
	// The ID of the device definition associated with this version.
	//
	// This value is a GUID.
	DeviceDefinitionId *string `field:"required" json:"deviceDefinitionId" yaml:"deviceDefinitionId"`
	// The devices in this version.
	Devices interface{} `field:"required" json:"devices" yaml:"devices"`
}

