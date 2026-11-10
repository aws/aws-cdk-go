package awssecurityagent


// The configuration for the private connection.
//
// Specify either a service-managed or a self-managed mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   modeProperty := &ModeProperty{
//   	SelfManaged: &SelfManagedProperty{
//   		ResourceConfigurationId: jsii.String("resourceConfigurationId"),
//   	},
//   	ServiceManaged: &ServiceManagedProperty{
//   		HostAddress: jsii.String("hostAddress"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcId: jsii.String("vpcId"),
//
//   		// the properties below are optional
//   		DnsResolution: jsii.String("dnsResolution"),
//   		IpAddressType: jsii.String("ipAddressType"),
//   		Ipv4AddressesPerEni: jsii.Number(123),
//   		PortRanges: []*string{
//   			jsii.String("portRanges"),
//   		},
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-mode.html
//
type CfnPrivateConnection_ModeProperty struct {
	// The configuration for a self-managed private connection, where you manage your own resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-mode.html#cfn-securityagent-privateconnection-mode-selfmanaged
	//
	SelfManaged interface{} `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// The configuration for a service-managed private connection, where the service manages the resource gateway lifecycle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityagent-privateconnection-mode.html#cfn-securityagent-privateconnection-mode-servicemanaged
	//
	ServiceManaged interface{} `field:"optional" json:"serviceManaged" yaml:"serviceManaged"`
}

