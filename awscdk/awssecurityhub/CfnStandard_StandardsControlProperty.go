package awssecurityhub


// Provides details about an individual security control.
//
// For a list of Security Hub controls, see [Security Hub controls reference](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-controls-reference.html) in the *AWS Security Hub User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   standardsControlProperty := &StandardsControlProperty{
//   	StandardsControlArn: jsii.String("standardsControlArn"),
//
//   	// the properties below are optional
//   	Reason: jsii.String("reason"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-standard-standardscontrol.html
//
type CfnStandard_StandardsControlProperty struct {
	// The Amazon Resource Name (ARN) of the control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-standard-standardscontrol.html#cfn-securityhub-standard-standardscontrol-standardscontrolarn
	//
	StandardsControlArn *string `field:"required" json:"standardsControlArn" yaml:"standardsControlArn"`
	// A user-defined reason for changing a control's enablement status in a specified standard.
	//
	// If you are disabling a control, then this property is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-standard-standardscontrol.html#cfn-securityhub-standard-standardscontrol-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
}

