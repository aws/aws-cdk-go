package awsdevopsagent

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnPrivateConnectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPrivateConnectionMixinProps := &CfnPrivateConnectionMixinProps{
//   	Certificate: jsii.String("certificate"),
//   	ConnectionConfiguration: &ConnectionConfigurationProperty{
//   		SelfManaged: &SelfManagedModeProperty{
//   			ResourceConfigurationId: jsii.String("resourceConfigurationId"),
//   		},
//   		ServiceManaged: &ServiceManagedModeProperty{
//   			HostAddress: jsii.String("hostAddress"),
//   			IpAddressType: jsii.String("ipAddressType"),
//   			Ipv4AddressesPerEni: jsii.Number(123),
//   			PortRanges: []*string{
//   				jsii.String("portRanges"),
//   			},
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   			VpcId: jsii.String("vpcId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html
//
type CfnPrivateConnectionMixinProps struct {
	// Certificate for the Private Connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html#cfn-devopsagent-privateconnection-certificate
	//
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// The connection configuration, either SelfManaged or ServiceManaged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html#cfn-devopsagent-privateconnection-connectionconfiguration
	//
	ConnectionConfiguration interface{} `field:"optional" json:"connectionConfiguration" yaml:"connectionConfiguration"`
	// Unique name for this Private Connection within the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html#cfn-devopsagent-privateconnection-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html#cfn-devopsagent-privateconnection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

