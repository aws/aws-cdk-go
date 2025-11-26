package previewawssecurityhubmixins


// Provides details about an individual security control.
//
// For a list of Security Hub controls, see [Security Hub controls reference](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-controls-reference.html) in the *Security Hub User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   standardsControlProperty := &StandardsControlProperty{
//   	Reason: jsii.String("reason"),
//   	StandardsControlArn: jsii.String("standardsControlArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-standard-standardscontrol.html
//
type CfnStandardPropsMixin_StandardsControlProperty struct {
	// A user-defined reason for changing a control's enablement status in a specified standard.
	//
	// If you are disabling a control, then this property is required.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-standard-standardscontrol.html#cfn-securityhub-standard-standardscontrol-reason
	//
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// The Amazon Resource Name (ARN) of the control.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-standard-standardscontrol.html#cfn-securityhub-standard-standardscontrol-standardscontrolarn
	//
	StandardsControlArn *string `field:"optional" json:"standardsControlArn" yaml:"standardsControlArn"`
}

