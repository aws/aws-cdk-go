package previewawsiotsitewisemixins


// Contains details for a gateway that runs on AWS IoT Greengrass V2 .
//
// To create a gateway that runs on AWS IoT Greengrass V2 , you must deploy the IoT SiteWise Edge component to your gateway device. Your [Greengrass device role](https://docs.aws.amazon.com/greengrass/v2/developerguide/device-service-role.html) must use the `AWSIoTSiteWiseEdgeAccess` policy. For more information, see [Using AWS IoT SiteWise at the edge](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/sw-gateways.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   greengrassV2Property := &GreengrassV2Property{
//   	CoreDeviceOperatingSystem: jsii.String("coreDeviceOperatingSystem"),
//   	CoreDeviceThingName: jsii.String("coreDeviceThingName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrassv2.html
//
type CfnGatewayPropsMixin_GreengrassV2Property struct {
	// The operating system of the core device in AWS IoT Greengrass V2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrassv2.html#cfn-iotsitewise-gateway-greengrassv2-coredeviceoperatingsystem
	//
	CoreDeviceOperatingSystem *string `field:"optional" json:"coreDeviceOperatingSystem" yaml:"coreDeviceOperatingSystem"`
	// The name of the AWS IoT thing for your AWS IoT Greengrass V2 core device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrassv2.html#cfn-iotsitewise-gateway-greengrassv2-coredevicethingname
	//
	CoreDeviceThingName *string `field:"optional" json:"coreDeviceThingName" yaml:"coreDeviceThingName"`
}

