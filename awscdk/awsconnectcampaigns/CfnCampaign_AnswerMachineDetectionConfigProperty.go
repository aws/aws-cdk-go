package awsconnectcampaigns


// Contains information about answering machine detection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   answerMachineDetectionConfigProperty := &AnswerMachineDetectionConfigProperty{
//   	EnableAnswerMachineDetection: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-answermachinedetectionconfig.html
//
type CfnCampaign_AnswerMachineDetectionConfigProperty struct {
	// Whether answering machine detection is enabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-answermachinedetectionconfig.html#cfn-connectcampaigns-campaign-answermachinedetectionconfig-enableanswermachinedetection
	//
	EnableAnswerMachineDetection interface{} `field:"required" json:"enableAnswerMachineDetection" yaml:"enableAnswerMachineDetection"`
}

