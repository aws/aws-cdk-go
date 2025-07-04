package awsiotsitewise


// Contains details for a AWS IoT SiteWise Edge gateway that runs on a Siemens Industrial Edge Device.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   siemensIEProperty := &SiemensIEProperty{
//   	IotCoreThingName: jsii.String("iotCoreThingName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-siemensie.html
//
type CfnGateway_SiemensIEProperty struct {
	// The name of the AWS IoT Thing for your AWS IoT SiteWise Edge gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-siemensie.html#cfn-iotsitewise-gateway-siemensie-iotcorethingname
	//
	IotCoreThingName *string `field:"required" json:"iotCoreThingName" yaml:"iotCoreThingName"`
}

