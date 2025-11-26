package previewawsec2mixins


// Specifies a Diffie-Hellman group number for the VPN tunnel for phase 2 IKE negotiations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   phase2DHGroupNumbersRequestListValueProperty := &Phase2DHGroupNumbersRequestListValueProperty{
//   	Value: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase2dhgroupnumbersrequestlistvalue.html
//
type CfnVPNConnectionPropsMixin_Phase2DHGroupNumbersRequestListValueProperty struct {
	// The Diffie-Hellmann group number.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase2dhgroupnumbersrequestlistvalue.html#cfn-ec2-vpnconnection-phase2dhgroupnumbersrequestlistvalue-value
	//
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

