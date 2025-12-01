package previewawsec2mixins


// Describes an exclusion configuration for VPC Encryption Control.
//
// For more information, see [Enforce VPC encryption in transit](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-encryption-controls.html) in the *Amazon VPC User Guide* .
//
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
	// The current state of the exclusion configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion.html#cfn-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// A message providing additional information about the exclusion state.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion.html#cfn-ec2-vpcencryptioncontrol-vpcencryptioncontrolexclusion-statemessage
	//
	StateMessage *string `field:"optional" json:"stateMessage" yaml:"stateMessage"`
}

