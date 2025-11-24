package mixinsawsmediaconnect


// The configuration settings for a public router network interface, including the list of allowed CIDR blocks.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   publicRouterNetworkInterfaceConfigurationProperty := &PublicRouterNetworkInterfaceConfigurationProperty{
//   	AllowRules: []interface{}{
//   		&PublicRouterNetworkInterfaceRuleProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routernetworkinterface-publicrouternetworkinterfaceconfiguration.html
//
type CfnRouterNetworkInterfacePropsMixin_PublicRouterNetworkInterfaceConfigurationProperty struct {
	// The list of allowed CIDR blocks for the public router network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routernetworkinterface-publicrouternetworkinterfaceconfiguration.html#cfn-mediaconnect-routernetworkinterface-publicrouternetworkinterfaceconfiguration-allowrules
	//
	AllowRules interface{} `field:"optional" json:"allowRules" yaml:"allowRules"`
}

