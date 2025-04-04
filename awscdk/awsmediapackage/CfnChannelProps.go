package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnChannel`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnChannelProps := &CfnChannelProps{
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EgressAccessLogs: &LogConfigurationProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   	},
//   	HlsIngest: &HlsIngestProperty{
//   		IngestEndpoints: []interface{}{
//   			&IngestEndpointProperty{
//   				Id: jsii.String("id"),
//   				Password: jsii.String("password"),
//   				Url: jsii.String("url"),
//   				Username: jsii.String("username"),
//   			},
//   		},
//   	},
//   	IngressAccessLogs: &LogConfigurationProperty{
//   		LogGroupName: jsii.String("logGroupName"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-channel.html
//
type CfnChannelProps struct {
	// Unique identifier that you assign to the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-channel.html#cfn-mediapackage-channel-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// Any descriptive information that you want to add to the channel for future identification purposes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-channel.html#cfn-mediapackage-channel-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configures egress access logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-channel.html#cfn-mediapackage-channel-egressaccesslogs
	//
	EgressAccessLogs interface{} `field:"optional" json:"egressAccessLogs" yaml:"egressAccessLogs"`
	// The input URL where the source stream should be sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-channel.html#cfn-mediapackage-channel-hlsingest
	//
	HlsIngest interface{} `field:"optional" json:"hlsIngest" yaml:"hlsIngest"`
	// Configures ingress access logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-channel.html#cfn-mediapackage-channel-ingressaccesslogs
	//
	IngressAccessLogs interface{} `field:"optional" json:"ingressAccessLogs" yaml:"ingressAccessLogs"`
	// The tags to assign to the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-channel.html#cfn-mediapackage-channel-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

