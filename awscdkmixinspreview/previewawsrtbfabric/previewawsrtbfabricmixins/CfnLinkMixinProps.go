package previewawsrtbfabricmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnLinkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLinkMixinProps := &CfnLinkMixinProps{
//   	GatewayId: jsii.String("gatewayId"),
//   	HttpResponderAllowed: jsii.Boolean(false),
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
//   	ModuleConfigurationList: []interface{}{
//   		&ModuleConfigurationProperty{
//   			DependsOn: []*string{
//   				jsii.String("dependsOn"),
//   			},
//   			ModuleParameters: &ModuleParametersProperty{
//   				NoBid: &NoBidModuleParametersProperty{
//   					PassThroughPercentage: jsii.Number(123),
//   					Reason: jsii.String("reason"),
//   					ReasonCode: jsii.Number(123),
//   				},
//   				OpenRtbAttribute: &OpenRtbAttributeModuleParametersProperty{
//   					Action: &ActionProperty{
//   						HeaderTag: &HeaderTagActionProperty{
//   							Name: jsii.String("name"),
//   							Value: jsii.String("value"),
//   						},
//   						NoBid: &NoBidActionProperty{
//   							NoBidReasonCode: jsii.Number(123),
//   						},
//   					},
//   					FilterConfiguration: []interface{}{
//   						&FilterProperty{
//   							Criteria: []interface{}{
//   								&FilterCriterionProperty{
//   									Path: jsii.String("path"),
//   									Values: []*string{
//   										jsii.String("values"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					FilterType: jsii.String("filterType"),
//   					HoldbackPercentage: jsii.Number(123),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Version: jsii.String("version"),
//   		},
//   	},
//   	PeerGatewayId: jsii.String("peerGatewayId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html
//
type CfnLinkMixinProps struct {
	// The unique identifier of the gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-gatewayid
	//
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Boolean to specify if an HTTP responder is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-httpresponderallowed
	//
	HttpResponderAllowed interface{} `field:"optional" json:"httpResponderAllowed" yaml:"httpResponderAllowed"`
	// Attributes of the link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-linkattributes
	//
	LinkAttributes interface{} `field:"optional" json:"linkAttributes" yaml:"linkAttributes"`
	// Settings for the application logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-linklogsettings
	//
	LinkLogSettings interface{} `field:"optional" json:"linkLogSettings" yaml:"linkLogSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-moduleconfigurationlist
	//
	ModuleConfigurationList interface{} `field:"optional" json:"moduleConfigurationList" yaml:"moduleConfigurationList"`
	// The unique identifier of the peer gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-peergatewayid
	//
	PeerGatewayId *string `field:"optional" json:"peerGatewayId" yaml:"peerGatewayId"`
	// A map of the key-value pairs of the tag or tags to assign to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

