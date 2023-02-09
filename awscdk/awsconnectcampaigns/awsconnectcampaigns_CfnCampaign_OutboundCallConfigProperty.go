package awsconnectcampaigns


// Contains outbound call configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outboundCallConfigProperty := &outboundCallConfigProperty{
//   	connectContactFlowArn: jsii.String("connectContactFlowArn"),
//   	connectQueueArn: jsii.String("connectQueueArn"),
//
//   	// the properties below are optional
//   	answerMachineDetectionConfig: &answerMachineDetectionConfigProperty{
//   		enableAnswerMachineDetection: jsii.Boolean(false),
//   	},
//   	connectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   }
//
type CfnCampaign_OutboundCallConfigProperty struct {
	// The Amazon Resource Name (ARN) of the flow.
	ConnectContactFlowArn *string `field:"required" json:"connectContactFlowArn" yaml:"connectContactFlowArn"`
	// The Amazon Resource Name (ARN) of the queue.
	ConnectQueueArn *string `field:"required" json:"connectQueueArn" yaml:"connectQueueArn"`
	// `CfnCampaign.OutboundCallConfigProperty.AnswerMachineDetectionConfig`.
	AnswerMachineDetectionConfig interface{} `field:"optional" json:"answerMachineDetectionConfig" yaml:"answerMachineDetectionConfig"`
	// The phone number associated with the outbound call.
	//
	// This is the caller ID that is displayed to customers when an agent calls them.
	ConnectSourcePhoneNumber *string `field:"optional" json:"connectSourcePhoneNumber" yaml:"connectSourcePhoneNumber"`
}

