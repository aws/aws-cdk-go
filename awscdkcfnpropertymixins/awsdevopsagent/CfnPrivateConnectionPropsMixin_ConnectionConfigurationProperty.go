package awsdevopsagent


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   connectionConfigurationProperty := &ConnectionConfigurationProperty{
//   	SelfManaged: &SelfManagedModeProperty{
//   		ResourceConfigurationId: jsii.String("resourceConfigurationId"),
//   	},
//   	ServiceManaged: &ServiceManagedModeProperty{
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-connectionconfiguration.html
//
type CfnPrivateConnectionPropsMixin_ConnectionConfigurationProperty struct {
	// Configuration for a self-managed Private Connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-connectionconfiguration.html#cfn-devopsagent-privateconnection-connectionconfiguration-selfmanaged
	//
	SelfManaged interface{} `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// Configuration for a service-managed Private Connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-privateconnection-connectionconfiguration.html#cfn-devopsagent-privateconnection-connectionconfiguration-servicemanaged
	//
	ServiceManaged interface{} `field:"optional" json:"serviceManaged" yaml:"serviceManaged"`
}

