package awsiotsitewise


// Contains details for a gateway that runs on AWS IoT Greengrass .
//
// To create a gateway that runs on AWS IoT Greengrass , you must add the IoT SiteWise connector to a Greengrass group and deploy it. Your Greengrass group must also have permissions to upload data to AWS IoT SiteWise . For more information, see [Ingesting data using a gateway](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/gateway-connector.html) in the *AWS IoT SiteWise User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   greengrassProperty := &GreengrassProperty{
//   	GroupArn: jsii.String("groupArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrass.html
//
type CfnGateway_GreengrassProperty struct {
	// The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the Greengrass group. For more information about how to find a group's ARN, see [ListGroups](https://docs.aws.amazon.com/greengrass/v1/apireference/listgroups-get.html) and [GetGroup](https://docs.aws.amazon.com/greengrass/v1/apireference/getgroup-get.html) in the *AWS IoT Greengrass V1 API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-gateway-greengrass.html#cfn-iotsitewise-gateway-greengrass-grouparn
	//
	GroupArn *string `field:"required" json:"groupArn" yaml:"groupArn"`
}

