package awsiotsitewise

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnGateway`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayProps := &cfnGatewayProps{
//   	gatewayName: jsii.String("gatewayName"),
//   	gatewayPlatform: &gatewayPlatformProperty{
//   		greengrass: &greengrassProperty{
//   			groupArn: jsii.String("groupArn"),
//   		},
//   		greengrassV2: &greengrassV2Property{
//   			coreDeviceThingName: jsii.String("coreDeviceThingName"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	gatewayCapabilitySummaries: []interface{}{
//   		&gatewayCapabilitySummaryProperty{
//   			capabilityNamespace: jsii.String("capabilityNamespace"),
//
//   			// the properties below are optional
//   			capabilityConfiguration: jsii.String("capabilityConfiguration"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnGatewayProps struct {
	// A unique, friendly name for the gateway.
	//
	// The maximum length is 256 characters with the pattern `[^\ u0000-\ u001F\ u007F]+` .
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// The gateway's platform.
	//
	// You can only specify one platform in a gateway.
	GatewayPlatform interface{} `field:"required" json:"gatewayPlatform" yaml:"gatewayPlatform"`
	// A list of gateway capability summaries that each contain a namespace and status.
	//
	// Each gateway capability defines data sources for the gateway. To retrieve a capability configuration's definition, use [DescribeGatewayCapabilityConfiguration](https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGatewayCapabilityConfiguration.html) .
	GatewayCapabilitySummaries interface{} `field:"optional" json:"gatewayCapabilitySummaries" yaml:"gatewayCapabilitySummaries"`
	// A list of key-value pairs that contain metadata for the gateway.
	//
	// For more information, see [Tagging your AWS IoT SiteWise resources](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/tag-resources.html) in the *AWS IoT SiteWise User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

