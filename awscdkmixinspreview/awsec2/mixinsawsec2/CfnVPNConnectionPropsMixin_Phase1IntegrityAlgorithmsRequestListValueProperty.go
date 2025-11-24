package mixinsawsec2


// Specifies the integrity algorithm for the VPN tunnel for phase 1 IKE negotiations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   phase1IntegrityAlgorithmsRequestListValueProperty := &Phase1IntegrityAlgorithmsRequestListValueProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1integrityalgorithmsrequestlistvalue.html
//
type CfnVPNConnectionPropsMixin_Phase1IntegrityAlgorithmsRequestListValueProperty struct {
	// The value for the integrity algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase1integrityalgorithmsrequestlistvalue.html#cfn-ec2-vpnconnection-phase1integrityalgorithmsrequestlistvalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

