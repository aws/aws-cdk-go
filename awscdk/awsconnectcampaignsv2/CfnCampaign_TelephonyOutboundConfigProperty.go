package awsconnectcampaignsv2


// The outbound configuration for telephony.
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
//   	RingTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html
//
type CfnCampaign_TelephonyOutboundConfigProperty struct {
	// The identifier of the published Amazon Connect contact flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-connectcontactflowid
	//
	ConnectContactFlowId *string `field:"required" json:"connectContactFlowId" yaml:"connectContactFlowId"`
	// The answering machine detection configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-answermachinedetectionconfig
	//
	AnswerMachineDetectionConfig interface{} `field:"optional" json:"answerMachineDetectionConfig" yaml:"answerMachineDetectionConfig"`
	// The Amazon Connect source phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-connectsourcephonenumber
	//
	ConnectSourcePhoneNumber *string `field:"optional" json:"connectSourcePhoneNumber" yaml:"connectSourcePhoneNumber"`
	// Maximum ring time for outbound calls in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-ringtimeout
	//
	RingTimeout *float64 `field:"optional" json:"ringTimeout" yaml:"ringTimeout"`
}

