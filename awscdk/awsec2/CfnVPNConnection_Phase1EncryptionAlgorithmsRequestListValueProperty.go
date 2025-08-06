package awsec2


// Specifies the encryption algorithm for the VPN tunnel for phase 1 IKE negotiations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   phase1EncryptionAlgorithmsRequestListValueProperty := &Phase1EncryptionAlgorithmsRequestListValueProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1encryptionalgorithmsrequestlistvalue.html
//
type CfnVPNConnection_Phase1EncryptionAlgorithmsRequestListValueProperty struct {
	// The value for the encryption algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1encryptionalgorithmsrequestlistvalue.html#cfn-ec2-vpnconnection-phase1encryptionalgorithmsrequestlistvalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

