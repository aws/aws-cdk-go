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
//   		AgentlessDialerConfig: &AgentlessDialerConfigProperty{
//   			DialingCapacity: jsii.Number(123),
//   		},
//   		PredictiveDialerConfig: &PredictiveDialerConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//
//   			// the properties below are optional
//   			DialingCapacity: jsii.Number(123),
//   		},
//   		ProgressiveDialerConfig: &ProgressiveDialerConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//
//   			// the properties below are optional
//   			DialingCapacity: jsii.Number(123),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	OutboundCallConfig: &OutboundCallConfigProperty{
//   		ConnectContactFlowArn: jsii.String("connectContactFlowArn"),
//
//   		// the properties below are optional
//   		AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   			EnableAnswerMachineDetection: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   		},
//   		ConnectQueueArn: jsii.String("connectQueueArn"),
//   		ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaigns-campaign.html
//
type CfnCampaignProps struct {
	// The Amazon Resource Name (ARN) of the Amazon Connect instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaigns-campaign.html#cfn-connectcampaigns-campaign-connectinstancearn
	//
	ConnectInstanceArn *string `field:"required" json:"connectInstanceArn" yaml:"connectInstanceArn"`
	// Contains information about the dialer configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaigns-campaign.html#cfn-connectcampaigns-campaign-dialerconfig
	//
	DialerConfig interface{} `field:"required" json:"dialerConfig" yaml:"dialerConfig"`
	// The name of the campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaigns-campaign.html#cfn-connectcampaigns-campaign-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains information about the outbound call configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaigns-campaign.html#cfn-connectcampaigns-campaign-outboundcallconfig
	//
	OutboundCallConfig interface{} `field:"required" json:"outboundCallConfig" yaml:"outboundCallConfig"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaigns-campaign.html#cfn-connectcampaigns-campaign-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

