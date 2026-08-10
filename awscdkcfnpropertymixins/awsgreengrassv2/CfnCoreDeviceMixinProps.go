package awsgreengrassv2


// Properties for CfnCoreDevicePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCoreDeviceMixinProps := &CfnCoreDeviceMixinProps{
//   	CoreDeviceThingName: jsii.String("coreDeviceThingName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-coredevice.html
//
type CfnCoreDeviceMixinProps struct {
	// The name of the core device.
	//
	// This is also the name of the IoT thing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-coredevice.html#cfn-greengrassv2-coredevice-coredevicethingname
	//
	CoreDeviceThingName *string `field:"optional" json:"coreDeviceThingName" yaml:"coreDeviceThingName"`
}

