package mixinsawsec2


// Specifies the encryption algorithm for the VPN tunnel for phase 2 IKE negotiations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   phase2EncryptionAlgorithmsRequestListValueProperty := &Phase2EncryptionAlgorithmsRequestListValueProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase2encryptionalgorithmsrequestlistvalue.html
//
type CfnVPNConnectionPropsMixin_Phase2EncryptionAlgorithmsRequestListValueProperty struct {
	// The encryption algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase2encryptionalgorithmsrequestlistvalue.html#cfn-ec2-vpnconnection-phase2encryptionalgorithmsrequestlistvalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

