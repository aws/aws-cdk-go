package awsconnectcampaignsv2


// Default Telephone Outbound config.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   telephonyOutboundConfigProperty := &TelephonyOutboundConfigProperty{
//   	ConnectContactFlowId: jsii.String("connectContactFlowId"),
//
//   	// the properties below are optional
//   	AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   		EnableAnswerMachineDetection: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   	},
//   	ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html
//
type CfnCampaign_TelephonyOutboundConfigProperty struct {
	// The identifier of the contact flow for the outbound call.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-connectcontactflowid
	//
	ConnectContactFlowId *string `field:"required" json:"connectContactFlowId" yaml:"connectContactFlowId"`
	// The configuration used for answering machine detection during outbound calls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-answermachinedetectionconfig
	//
	AnswerMachineDetectionConfig interface{} `field:"optional" json:"answerMachineDetectionConfig" yaml:"answerMachineDetectionConfig"`
	// The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-connectsourcephonenumber
	//
	ConnectSourcePhoneNumber *string `field:"optional" json:"connectSourcePhoneNumber" yaml:"connectSourcePhoneNumber"`
}

