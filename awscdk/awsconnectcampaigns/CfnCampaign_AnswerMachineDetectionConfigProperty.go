package awsconnectcampaigns


// The configuration used for answering machine detection during outbound calls.
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
	// Flag to decided whether outbound calls should have answering machine detection enabled or not.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaigns-campaign-answermachinedetectionconfig.html#cfn-connectcampaigns-campaign-answermachinedetectionconfig-enableanswermachinedetection
	//
	EnableAnswerMachineDetection interface{} `field:"required" json:"enableAnswerMachineDetection" yaml:"enableAnswerMachineDetection"`
}

