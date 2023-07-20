package awsiotsitewise


// Contains a gateway's platform information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayPlatformProperty := &GatewayPlatformProperty{
//   	Greengrass: &GreengrassProperty{
//   		GroupArn: jsii.String("groupArn"),
//   	},
//   	GreengrassV2: &GreengrassV2Property{
//   		CoreDeviceThingName: jsii.String("coreDeviceThingName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewayplatform.html
//
type CfnGateway_GatewayPlatformProperty struct {
	// A gateway that runs on AWS IoT Greengrass .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewayplatform.html#cfn-iotsitewise-gateway-gatewayplatform-greengrass
	//
	Greengrass interface{} `field:"optional" json:"greengrass" yaml:"greengrass"`
	// A gateway that runs on AWS IoT Greengrass V2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-gatewayplatform.html#cfn-iotsitewise-gateway-gatewayplatform-greengrassv2
	//
	GreengrassV2 interface{} `field:"optional" json:"greengrassV2" yaml:"greengrassV2"`
}

