package mixinsawsec2


// The IKE version that is permitted for the VPN tunnel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   iKEVersionsRequestListValueProperty := map[string]*string{
//   	"value": jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-ikeversionsrequestlistvalue.html
//
type CfnVPNConnectionPropsMixin_IKEVersionsRequestListValueProperty struct {
	// The IKE version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpnconnection-ikeversionsrequestlistvalue.html#cfn-ec2-vpnconnection-ikeversionsrequestlistvalue-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

