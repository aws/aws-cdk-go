package previewawsiotmixins


// A structure containing the alert target ARN and the role ARN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   alertTargetProperty := &AlertTargetProperty{
//   	AlertTargetArn: jsii.String("alertTargetArn"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-alerttarget.html
//
type CfnSecurityProfilePropsMixin_AlertTargetProperty struct {
	// The Amazon Resource Name (ARN) of the notification target to which alerts are sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-alerttarget.html#cfn-iot-securityprofile-alerttarget-alerttargetarn
	//
	AlertTargetArn *string `field:"optional" json:"alertTargetArn" yaml:"alertTargetArn"`
	// The ARN of the role that grants permission to send alerts to the notification target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-securityprofile-alerttarget.html#cfn-iot-securityprofile-alerttarget-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

