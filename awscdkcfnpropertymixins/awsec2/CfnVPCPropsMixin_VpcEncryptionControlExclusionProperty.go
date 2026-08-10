package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vpcEncryptionControlExclusionProperty := &VpcEncryptionControlExclusionProperty{
//   	State: jsii.String("state"),
//   	StateMessage: jsii.String("stateMessage"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusion.html
//
type CfnVPCPropsMixin_VpcEncryptionControlExclusionProperty struct {
	// The exclusion state of the resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusion.html#cfn-ec2-vpc-vpcencryptioncontrolexclusion-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
	// A message describing the exclusion state of the resource type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpc-vpcencryptioncontrolexclusion.html#cfn-ec2-vpc-vpcencryptioncontrolexclusion-statemessage
	//
	StateMessage *string `field:"optional" json:"stateMessage" yaml:"stateMessage"`
}

