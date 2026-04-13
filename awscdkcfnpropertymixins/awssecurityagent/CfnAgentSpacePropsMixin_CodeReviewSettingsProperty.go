package awssecurityagent


// Details of code review settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   codeReviewSettingsProperty := &CodeReviewSettingsProperty{
//   	ControlsScanning: jsii.Boolean(false),
//   	GeneralPurposeScanning: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-codereviewsettings.html
//
type CfnAgentSpacePropsMixin_CodeReviewSettingsProperty struct {
	// Whether Controls are utilized for code review analysis.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-codereviewsettings.html#cfn-securityagent-agentspace-codereviewsettings-controlsscanning
	//
	ControlsScanning interface{} `field:"optional" json:"controlsScanning" yaml:"controlsScanning"`
	// Whether general purpose analysis is performed for code review.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-agentspace-codereviewsettings.html#cfn-securityagent-agentspace-codereviewsettings-generalpurposescanning
	//
	GeneralPurposeScanning interface{} `field:"optional" json:"generalPurposeScanning" yaml:"generalPurposeScanning"`
}

