package mixinsawsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routerNetworkInterfaceConfigurationProperty := &RouterNetworkInterfaceConfigurationProperty{
//   	Public: &PublicRouterNetworkInterfaceConfigurationProperty{
//   		AllowRules: []interface{}{
//   			&PublicRouterNetworkInterfaceRuleProperty{
//   				Cidr: jsii.String("cidr"),
//   			},
//   		},
//   	},
//   	Vpc: &VpcRouterNetworkInterfaceConfigurationProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetId: jsii.String("subnetId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routernetworkinterface-routernetworkinterfaceconfiguration.html
//
type CfnRouterNetworkInterfacePropsMixin_RouterNetworkInterfaceConfigurationProperty struct {
	// The configuration settings for a public router network interface, including the list of allowed CIDR blocks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routernetworkinterface-routernetworkinterfaceconfiguration.html#cfn-mediaconnect-routernetworkinterface-routernetworkinterfaceconfiguration-public
	//
	Public interface{} `field:"optional" json:"public" yaml:"public"`
	// The configuration settings for a router network interface within a VPC, including the security group IDs and subnet ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-routernetworkinterface-routernetworkinterfaceconfiguration.html#cfn-mediaconnect-routernetworkinterface-routernetworkinterfaceconfiguration-vpc
	//
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

