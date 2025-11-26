package previewawsec2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcEncryptionControlExclusionProperty := &VpcEncryptionControlExclusionProperty{
//   	State: jsii.String("state"),
//   	StateMessage: jsii.String("stateMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion.html
//
type CfnVPCEncryptionControlPropsMixin_VpcEncryptionControlExclusionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion.html#cfn-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion.html#cfn-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion-statemessage
	//
	StateMessage *string `field:"optional" json:"stateMessage" yaml:"stateMessage"`
}

