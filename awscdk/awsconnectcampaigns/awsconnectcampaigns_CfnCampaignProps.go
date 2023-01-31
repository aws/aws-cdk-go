package awsconnectcampaigns

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCampaign`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCampaignProps := &cfnCampaignProps{
//   	connectInstanceArn: jsii.String("connectInstanceArn"),
//   	dialerConfig: &dialerConfigProperty{
//   		predictiveDialerConfig: &predictiveDialerConfigProperty{
//   			bandwidthAllocation: jsii.Number(123),
//   		},
//   		progressiveDialerConfig: &progressiveDialerConfigProperty{
//   			bandwidthAllocation: jsii.Number(123),
//   		},
//   	},
//   	name: jsii.String("name"),
//   	outboundCallConfig: &outboundCallConfigProperty{
//   		connectContactFlowArn: jsii.String("connectContactFlowArn"),
//   		connectQueueArn: jsii.String("connectQueueArn"),
//
//   		// the properties below are optional
//   		connectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   	},
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnCampaignProps struct {
	// The Amazon Resource Name (ARN) of the Amazon Connect instance.
	ConnectInstanceArn *string `field:"required" json:"connectInstanceArn" yaml:"connectInstanceArn"`
	// Contains information about the dialer configuration.
	DialerConfig interface{} `field:"required" json:"dialerConfig" yaml:"dialerConfig"`
	// The name of the campaign.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains information about the outbound call configuration.
	OutboundCallConfig interface{} `field:"required" json:"outboundCallConfig" yaml:"outboundCallConfig"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

