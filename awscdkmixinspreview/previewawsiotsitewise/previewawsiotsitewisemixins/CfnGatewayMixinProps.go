package previewawsiotsitewisemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnGatewayPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnGatewayMixinProps := &CfnGatewayMixinProps{
//   	GatewayCapabilitySummaries: []interface{}{
//   		&GatewayCapabilitySummaryProperty{
//   			CapabilityConfiguration: jsii.String("capabilityConfiguration"),
//   			CapabilityNamespace: jsii.String("capabilityNamespace"),
//   		},
//   	},
//   	GatewayName: jsii.String("gatewayName"),
//   	GatewayPlatform: &GatewayPlatformProperty{
//   		Greengrass: &GreengrassProperty{
//   			GroupArn: jsii.String("groupArn"),
//   		},
//   		GreengrassV2: &GreengrassV2Property{
//   			CoreDeviceOperatingSystem: jsii.String("coreDeviceOperatingSystem"),
//   			CoreDeviceThingName: jsii.String("coreDeviceThingName"),
//   		},
//   		SiemensIe: &SiemensIEProperty{
//   			IotCoreThingName: jsii.String("iotCoreThingName"),
//   		},
//   	},
//   	GatewayVersion: jsii.String("gatewayVersion"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html
//
type CfnGatewayMixinProps struct {
	// A list of gateway capability summaries that each contain a namespace and status.
	//
	// Each gateway capability defines data sources for the gateway. To retrieve a capability configuration's definition, use [DescribeGatewayCapabilityConfiguration](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGatewayCapabilityConfiguration.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewaycapabilitysummaries
	//
	GatewayCapabilitySummaries interface{} `field:"optional" json:"gatewayCapabilitySummaries" yaml:"gatewayCapabilitySummaries"`
	// A unique name for the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayname
	//
	GatewayName *string `field:"optional" json:"gatewayName" yaml:"gatewayName"`
	// The gateway's platform.
	//
	// You can only specify one platform in a gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayplatform
	//
	GatewayPlatform interface{} `field:"optional" json:"gatewayPlatform" yaml:"gatewayPlatform"`
	// The version of the gateway.
	//
	// A value of `3` indicates an MQTT-enabled, V3 gateway, while `2` indicates a Classic streams, V2 gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-gatewayversion
	//
	GatewayVersion *string `field:"optional" json:"gatewayVersion" yaml:"gatewayVersion"`
	// A list of key-value pairs that contain metadata for the gateway.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-gateway.html#cfn-iotsitewise-gateway-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

