package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayProps := &CfnGatewayProps{
//   	GatewayName: jsii.String("gatewayName"),
//   	GatewayPlatform: &GatewayPlatformProperty{
//   		Greengrass: &GreengrassProperty{
//   			GroupArn: jsii.String("groupArn"),
//   		},
//   		GreengrassV2: &GreengrassV2Property{
//   			CoreDeviceThingName: jsii.String("coreDeviceThingName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	GatewayCapabilitySummaries: []interface{}{
//   		&GatewayCapabilitySummaryProperty{
//   			CapabilityNamespace: jsii.String("capabilityNamespace"),
//
//   			// the properties below are optional
//   			CapabilityConfiguration: jsii.String("capabilityConfiguration"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html
//
type CfnGatewayProps struct {
	// A unique, friendly name for the gateway.
	//
	// The maximum length is 256 characters with the pattern `[^\u0000-\u001F\u007F]+` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayname
	//
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// The gateway's platform.
	//
	// You can only specify one platform in a gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayplatform
	//
	GatewayPlatform interface{} `field:"required" json:"gatewayPlatform" yaml:"gatewayPlatform"`
	// A list of gateway capability summaries that each contain a namespace and status.
	//
	// Each gateway capability defines data sources for the gateway. To retrieve a capability configuration's definition, use [DescribeGatewayCapabilityConfiguration](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGatewayCapabilityConfiguration.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewaycapabilitysummaries
	//
	GatewayCapabilitySummaries interface{} `field:"optional" json:"gatewayCapabilitySummaries" yaml:"gatewayCapabilitySummaries"`
	// A list of key-value pairs that contain metadata for the gateway.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

