package awsconnectcampaigns


// Contains outbound call configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outboundCallConfigProperty := &OutboundCallConfigProperty{
//   	ConnectContactFlowArn: jsii.String("connectContactFlowArn"),
//   	ConnectQueueArn: jsii.String("connectQueueArn"),
//
//   	// the properties below are optional
//   	AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   		EnableAnswerMachineDetection: jsii.Boolean(false),
//   	},
//   	ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-outboundcallconfig.html
//
type CfnCampaign_OutboundCallConfigProperty struct {
	// The Amazon Resource Name (ARN) of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-outboundcallconfig.html#cfn-connectcampaigns-campaign-outboundcallconfig-connectcontactflowarn
	//
	ConnectContactFlowArn *string `field:"required" json:"connectContactFlowArn" yaml:"connectContactFlowArn"`
	// The Amazon Resource Name (ARN) of the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-outboundcallconfig.html#cfn-connectcampaigns-campaign-outboundcallconfig-connectqueuearn
	//
	ConnectQueueArn *string `field:"required" json:"connectQueueArn" yaml:"connectQueueArn"`
	// The configuration used for answering machine detection during outbound calls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-outboundcallconfig.html#cfn-connectcampaigns-campaign-outboundcallconfig-answermachinedetectionconfig
	//
	AnswerMachineDetectionConfig interface{} `field:"optional" json:"answerMachineDetectionConfig" yaml:"answerMachineDetectionConfig"`
	// The phone number associated with the outbound call.
	//
	// This is the caller ID that is displayed to customers when an agent calls them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-outboundcallconfig.html#cfn-connectcampaigns-campaign-outboundcallconfig-connectsourcephonenumber
	//
	ConnectSourcePhoneNumber *string `field:"optional" json:"connectSourcePhoneNumber" yaml:"connectSourcePhoneNumber"`
}

