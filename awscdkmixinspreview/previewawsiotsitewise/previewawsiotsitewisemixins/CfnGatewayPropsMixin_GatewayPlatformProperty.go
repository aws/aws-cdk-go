package previewawsiotsitewisemixins


// The gateway's platform configuration. You can only specify one platform type in a gateway.
//
// (Legacy only) For Greengrass V1 gateways, specify the `greengrass` parameter with a valid Greengrass group ARN.
//
// For Greengrass V2 gateways, specify the `greengrassV2` parameter with a valid core device thing name. If creating a V3 gateway ( `gatewayVersion=3` ), you must also specify the `coreDeviceOperatingSystem` .
//
// For Siemens Industrial Edge gateways, specify the `siemensIE` parameter with a valid IoT Core thing name.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gatewayPlatformProperty := &GatewayPlatformProperty{
//   	Greengrass: &GreengrassProperty{
//   		GroupArn: jsii.String("groupArn"),
//   	},
//   	GreengrassV2: &GreengrassV2Property{
//   		CoreDeviceOperatingSystem: jsii.String("coreDeviceOperatingSystem"),
//   		CoreDeviceThingName: jsii.String("coreDeviceThingName"),
//   	},
//   	SiemensIe: &SiemensIEProperty{
//   		IotCoreThingName: jsii.String("iotCoreThingName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewayplatform.html
//
type CfnGatewayPropsMixin_GatewayPlatformProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewayplatform.html#cfn-iotsitewise-gateway-gatewayplatform-greengrass
	//
	Greengrass interface{} `field:"optional" json:"greengrass" yaml:"greengrass"`
	// A gateway that runs on AWS IoT Greengrass V2 .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewayplatform.html#cfn-iotsitewise-gateway-gatewayplatform-greengrassv2
	//
	GreengrassV2 interface{} `field:"optional" json:"greengrassV2" yaml:"greengrassV2"`
	// An AWS IoT SiteWise Edge gateway that runs on a Siemens Industrial Edge Device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewayplatform.html#cfn-iotsitewise-gateway-gatewayplatform-siemensie
	//
	SiemensIe interface{} `field:"optional" json:"siemensIe" yaml:"siemensIe"`
}

