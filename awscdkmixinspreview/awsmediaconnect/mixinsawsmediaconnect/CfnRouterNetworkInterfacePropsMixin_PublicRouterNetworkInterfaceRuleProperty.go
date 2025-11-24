package mixinsawsmediaconnect


// A rule that allows a specific CIDR block to access the public router network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   publicRouterNetworkInterfaceRuleProperty := &PublicRouterNetworkInterfaceRuleProperty{
//   	Cidr: jsii.String("cidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routernetworkinterface-publicrouternetworkinterfacerule.html
//
type CfnRouterNetworkInterfacePropsMixin_PublicRouterNetworkInterfaceRuleProperty struct {
	// The CIDR block that is allowed to access the public router network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routernetworkinterface-publicrouternetworkinterfacerule.html#cfn-mediaconnect-routernetworkinterface-publicrouternetworkinterfacerule-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

