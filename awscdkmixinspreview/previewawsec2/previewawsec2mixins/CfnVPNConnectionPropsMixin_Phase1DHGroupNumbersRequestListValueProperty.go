package previewawsec2mixins


// Specifies a Diffie-Hellman group number for the VPN tunnel for phase 1 IKE negotiations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   phase1DHGroupNumbersRequestListValueProperty := &Phase1DHGroupNumbersRequestListValueProperty{
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1dhgroupnumbersrequestlistvalue.html
//
type CfnVPNConnectionPropsMixin_Phase1DHGroupNumbersRequestListValueProperty struct {
	// The Diffie-Hellmann group number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1dhgroupnumbersrequestlistvalue.html#cfn-ec2-vpnconnection-phase1dhgroupnumbersrequestlistvalue-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

