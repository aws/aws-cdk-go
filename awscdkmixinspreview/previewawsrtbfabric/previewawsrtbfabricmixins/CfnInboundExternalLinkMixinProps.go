package previewawsrtbfabricmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnInboundExternalLinkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInboundExternalLinkMixinProps := &CfnInboundExternalLinkMixinProps{
//   	GatewayId: jsii.String("gatewayId"),
//   	LinkAttributes: &LinkAttributesProperty{
//   		CustomerProvidedId: jsii.String("customerProvidedId"),
//   		ResponderErrorMasking: []interface{}{
//   			&ResponderErrorMaskingForHttpCodeProperty{
//   				Action: jsii.String("action"),
//   				HttpCode: jsii.String("httpCode"),
//   				LoggingTypes: []*string{
//   					jsii.String("loggingTypes"),
//   				},
//   				ResponseLoggingPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	LinkLogSettings: &LinkLogSettingsProperty{
//   		ApplicationLogs: &ApplicationLogsProperty{
//   			LinkApplicationLogSampling: &LinkApplicationLogSamplingProperty{
//   				ErrorLog: jsii.Number(123),
//   				FilterLog: jsii.Number(123),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-inboundexternallink.html
//
type CfnInboundExternalLinkMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-inboundexternallink.html#cfn-rtbfabric-inboundexternallink-gatewayid
	//
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-inboundexternallink.html#cfn-rtbfabric-inboundexternallink-linkattributes
	//
	LinkAttributes interface{} `field:"optional" json:"linkAttributes" yaml:"linkAttributes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-inboundexternallink.html#cfn-rtbfabric-inboundexternallink-linklogsettings
	//
	LinkLogSettings interface{} `field:"optional" json:"linkLogSettings" yaml:"linkLogSettings"`
	// Tags to assign to the Link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-inboundexternallink.html#cfn-rtbfabric-inboundexternallink-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

