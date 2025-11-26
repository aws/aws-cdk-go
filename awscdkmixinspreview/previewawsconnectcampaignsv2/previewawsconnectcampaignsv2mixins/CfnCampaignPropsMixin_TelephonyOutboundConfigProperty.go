package previewawsconnectcampaignsv2mixins


// The outbound configuration for telephony.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   telephonyOutboundConfigProperty := &TelephonyOutboundConfigProperty{
//   	AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   		AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   		EnableAnswerMachineDetection: jsii.Boolean(false),
//   	},
//   	ConnectContactFlowId: jsii.String("connectContactFlowId"),
//   	ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   	RingTimeout: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html
//
type CfnCampaignPropsMixin_TelephonyOutboundConfigProperty struct {
	// The answering machine detection configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-answermachinedetectionconfig
	//
	AnswerMachineDetectionConfig interface{} `field:"optional" json:"answerMachineDetectionConfig" yaml:"answerMachineDetectionConfig"`
	// The identifier of the published Amazon Connect contact flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-connectcontactflowid
	//
	ConnectContactFlowId *string `field:"optional" json:"connectContactFlowId" yaml:"connectContactFlowId"`
	// The Amazon Connect source phone number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-connectsourcephonenumber
	//
	ConnectSourcePhoneNumber *string `field:"optional" json:"connectSourcePhoneNumber" yaml:"connectSourcePhoneNumber"`
	// Maximum ring time for outbound calls in seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonyoutboundconfig.html#cfn-connectcampaignsv2-campaign-telephonyoutboundconfig-ringtimeout
	//
	RingTimeout *float64 `field:"optional" json:"ringTimeout" yaml:"ringTimeout"`
}

