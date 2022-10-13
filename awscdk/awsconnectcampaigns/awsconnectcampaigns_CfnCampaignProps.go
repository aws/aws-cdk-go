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
	// `AWS::ConnectCampaigns::Campaign.ConnectInstanceArn`.
	ConnectInstanceArn *string `field:"required" json:"connectInstanceArn" yaml:"connectInstanceArn"`
	// `AWS::ConnectCampaigns::Campaign.DialerConfig`.
	DialerConfig interface{} `field:"required" json:"dialerConfig" yaml:"dialerConfig"`
	// `AWS::ConnectCampaigns::Campaign.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::ConnectCampaigns::Campaign.OutboundCallConfig`.
	OutboundCallConfig interface{} `field:"required" json:"outboundCallConfig" yaml:"outboundCallConfig"`
	// `AWS::ConnectCampaigns::Campaign.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

