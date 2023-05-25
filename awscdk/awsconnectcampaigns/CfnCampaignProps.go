package awsconnectcampaigns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCampaign`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCampaignProps := &CfnCampaignProps{
//   	ConnectInstanceArn: jsii.String("connectInstanceArn"),
//   	DialerConfig: &DialerConfigProperty{
//   		PredictiveDialerConfig: &PredictiveDialerConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//   		},
//   		ProgressiveDialerConfig: &ProgressiveDialerConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OutboundCallConfig: &OutboundCallConfigProperty{
//   		ConnectContactFlowArn: jsii.String("connectContactFlowArn"),
//   		ConnectQueueArn: jsii.String("connectQueueArn"),
//
//   		// the properties below are optional
//   		AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   			EnableAnswerMachineDetection: jsii.Boolean(false),
//   		},
//   		ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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

