package previewawsconnectcampaignsv2mixins


// Contains answering machine detection configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   answerMachineDetectionConfigProperty := &AnswerMachineDetectionConfigProperty{
//   	AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   	EnableAnswerMachineDetection: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-answermachinedetectionconfig.html
//
type CfnCampaignPropsMixin_AnswerMachineDetectionConfigProperty struct {
	// Whether or not waiting for an answer machine prompt is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-answermachinedetectionconfig.html#cfn-connectcampaignsv2-campaign-answermachinedetectionconfig-awaitanswermachineprompt
	//
	AwaitAnswerMachinePrompt interface{} `field:"optional" json:"awaitAnswerMachinePrompt" yaml:"awaitAnswerMachinePrompt"`
	// Enables answering machine detection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-answermachinedetectionconfig.html#cfn-connectcampaignsv2-campaign-answermachinedetectionconfig-enableanswermachinedetection
	//
	EnableAnswerMachineDetection interface{} `field:"optional" json:"enableAnswerMachineDetection" yaml:"enableAnswerMachineDetection"`
}

