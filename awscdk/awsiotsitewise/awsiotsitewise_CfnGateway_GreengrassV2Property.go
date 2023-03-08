package awsiotsitewise


// Contains details for a gateway that runs on AWS IoT Greengrass V2.
//
// To create a gateway that runs on AWS IoT Greengrass V2, you must deploy the IoT SiteWise Edge component to your gateway device. Your [Greengrass device role](https://docs.aws.amazon.com/greengrass/v2/developerguide/device-service-role.html) must use the `AWSIoTSiteWiseEdgeAccess` policy. For more information, see [Using AWS IoT SiteWise at the edge](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/sw-gateways.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   greengrassV2Property := &greengrassV2Property{
//   	coreDeviceThingName: jsii.String("coreDeviceThingName"),
//   }
//
type CfnGateway_GreengrassV2Property struct {
	// The name of the AWS IoT thing for your AWS IoT Greengrass V2 core device.
	CoreDeviceThingName *string `field:"required" json:"coreDeviceThingName" yaml:"coreDeviceThingName"`
}

