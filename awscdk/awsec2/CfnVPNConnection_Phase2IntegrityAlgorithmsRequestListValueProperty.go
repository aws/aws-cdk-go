package awsec2


// Specifies the integrity algorithm for the VPN tunnel for phase 2 IKE negotiations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   phase2IntegrityAlgorithmsRequestListValueProperty := &Phase2IntegrityAlgorithmsRequestListValueProperty{
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase2integrityalgorithmsrequestlistvalue.html
//
type CfnVPNConnection_Phase2IntegrityAlgorithmsRequestListValueProperty struct {
	// The integrity algorithm.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-phase2integrityalgorithmsrequestlistvalue.html#cfn-ec2-vpnconnection-phase2integrityalgorithmsrequestlistvalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

