package awssecurityagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPrivateConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPrivateConnectionProps := &CfnPrivateConnectionProps{
//   	PrivateConnectionName: jsii.String("privateConnectionName"),
//
//   	// the properties below are optional
//   	Certificate: jsii.String("certificate"),
//   	Mode: &ModeProperty{
//   		SelfManaged: &SelfManagedProperty{
//   			ResourceConfigurationId: jsii.String("resourceConfigurationId"),
//   		},
//   		ServiceManaged: &ServiceManagedProperty{
//   			HostAddress: jsii.String("hostAddress"),
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   			VpcId: jsii.String("vpcId"),
//
//   			// the properties below are optional
//   			DnsResolution: jsii.String("dnsResolution"),
//   			IpAddressType: jsii.String("ipAddressType"),
//   			Ipv4AddressesPerEni: jsii.Number(123),
//   			PortRanges: []*string{
//   				jsii.String("portRanges"),
//   			},
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-privateconnection.html
//
type CfnPrivateConnectionProps struct {
	// A unique name for the private connection within your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-privateconnection.html#cfn-securityagent-privateconnection-privateconnectionname
	//
	PrivateConnectionName *string `field:"required" json:"privateConnectionName" yaml:"privateConnectionName"`
	// The PEM-encoded certificate chain for the private connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-privateconnection.html#cfn-securityagent-privateconnection-certificate
	//
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The configuration for the private connection.
	//
	// Specify either a service-managed or a self-managed mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-privateconnection.html#cfn-securityagent-privateconnection-mode
	//
	Mode interface{} `field:"optional" json:"mode" yaml:"mode"`
	// The tags attached to the private connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityagent-privateconnection.html#cfn-securityagent-privateconnection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

