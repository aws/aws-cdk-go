package previewawsconnectcampaignsmixins


// Contains information about answering machine detection.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-answermachinedetectionconfig.html
//
type CfnCampaignPropsMixin_AnswerMachineDetectionConfigProperty struct {
	// Whether waiting for answer machine prompt is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-answermachinedetectionconfig.html#cfn-connectcampaigns-campaign-answermachinedetectionconfig-awaitanswermachineprompt
	//
	AwaitAnswerMachinePrompt interface{} `field:"optional" json:"awaitAnswerMachinePrompt" yaml:"awaitAnswerMachinePrompt"`
	// Whether answering machine detection is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-answermachinedetectionconfig.html#cfn-connectcampaigns-campaign-answermachinedetectionconfig-enableanswermachinedetection
	//
	EnableAnswerMachineDetection interface{} `field:"optional" json:"enableAnswerMachineDetection" yaml:"enableAnswerMachineDetection"`
}

