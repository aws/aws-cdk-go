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
type CfnCampaign_AnswerMachineDetectionConfigProperty struct {
	// Whether answering machine detection is enabled.
	EnableAnswerMachineDetection interface{} `field:"required" json:"enableAnswerMachineDetection" yaml:"enableAnswerMachineDetection"`
}

