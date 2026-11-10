package awssecurityagent


// The configuration for the private connection.
//
// Specify either a service-managed or a self-managed mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   modeProperty := &ModeProperty{
//   	SelfManaged: &SelfManagedProperty{
//   		ResourceConfigurationId: jsii.String("resourceConfigurationId"),
//   	},
//   	ServiceManaged: &ServiceManagedProperty{
//   		DnsResolution: jsii.String("dnsResolution"),
//   		HostAddress: jsii.String("hostAddress"),
//   		IpAddressType: jsii.String("ipAddressType"),
//   		Ipv4AddressesPerEni: jsii.Number(123),
//   		PortRanges: []*string{
//   			jsii.String("portRanges"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-mode.html
//
type CfnPrivateConnectionPropsMixin_ModeProperty struct {
	// The configuration for a self-managed private connection, where you manage your own resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-mode.html#cfn-securityagent-privateconnection-mode-selfmanaged
	//
	SelfManaged interface{} `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// The configuration for a service-managed private connection, where the service manages the resource gateway lifecycle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-mode.html#cfn-securityagent-privateconnection-mode-servicemanaged
	//
	ServiceManaged interface{} `field:"optional" json:"serviceManaged" yaml:"serviceManaged"`
}

