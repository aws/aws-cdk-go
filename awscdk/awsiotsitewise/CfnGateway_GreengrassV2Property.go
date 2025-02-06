package awsiotsitewise


// Contains details for a gateway that runs on AWS IoT Greengrass V2 .
//
// To create a gateway that runs on AWS IoT Greengrass V2 , you must deploy the IoT SiteWise Edge component to your gateway device. Your [Greengrass device role](https://docs.aws.amazon.com/greengrass/v2/developerguide/device-service-role.html) must use the `AWSIoTSiteWiseEdgeAccess` policy. For more information, see [Using AWS IoT SiteWise at the edge](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/sw-gateways.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   greengrassV2Property := &GreengrassV2Property{
//   	CoreDeviceThingName: jsii.String("coreDeviceThingName"),
//
//   	// the properties below are optional
//   	CoreDeviceOperatingSystem: jsii.String("coreDeviceOperatingSystem"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrassv2.html
//
type CfnGateway_GreengrassV2Property struct {
	// The name of the AWS IoT thing for your AWS IoT Greengrass V2 core device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrassv2.html#cfn-iotsitewise-gateway-greengrassv2-coredevicethingname
	//
	CoreDeviceThingName *string `field:"required" json:"coreDeviceThingName" yaml:"coreDeviceThingName"`
	// The operating system of the core device in AWS IoT Greengrass V2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrassv2.html#cfn-iotsitewise-gateway-greengrassv2-coredeviceoperatingsystem
	//
	CoreDeviceOperatingSystem *string `field:"optional" json:"coreDeviceOperatingSystem" yaml:"coreDeviceOperatingSystem"`
}

