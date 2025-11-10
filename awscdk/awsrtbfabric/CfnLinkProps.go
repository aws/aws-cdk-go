package awsrtbfabric

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLink`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLinkProps := &CfnLinkProps{
//   	GatewayId: jsii.String("gatewayId"),
//   	LinkLogSettings: &LinkLogSettingsProperty{
//   		ApplicationLogs: &ApplicationLogsProperty{
//   			LinkApplicationLogSampling: &LinkApplicationLogSamplingProperty{
//   				ErrorLog: jsii.Number(123),
//   				FilterLog: jsii.Number(123),
//   			},
//   		},
//   	},
//   	PeerGatewayId: jsii.String("peerGatewayId"),
//
//   	// the properties below are optional
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
//
//   				// the properties below are optional
//   				ResponseLoggingPercentage: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ModuleConfigurationList: []interface{}{
//   		&ModuleConfigurationProperty{
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
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
//   			Version: jsii.String("version"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html
//
type CfnLinkProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-gatewayid
	//
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-linklogsettings
	//
	LinkLogSettings interface{} `field:"required" json:"linkLogSettings" yaml:"linkLogSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-peergatewayid
	//
	PeerGatewayId *string `field:"required" json:"peerGatewayId" yaml:"peerGatewayId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-httpresponderallowed
	//
	HttpResponderAllowed interface{} `field:"optional" json:"httpResponderAllowed" yaml:"httpResponderAllowed"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-linkattributes
	//
	LinkAttributes interface{} `field:"optional" json:"linkAttributes" yaml:"linkAttributes"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-moduleconfigurationlist
	//
	ModuleConfigurationList interface{} `field:"optional" json:"moduleConfigurationList" yaml:"moduleConfigurationList"`
	// Tags to assign to the Link.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rtbfabric-link.html#cfn-rtbfabric-link-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

