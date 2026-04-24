package awsdevopsagent

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
//   	ConnectionConfiguration: &ConnectionConfigurationProperty{
//   		SelfManaged: &SelfManagedModeProperty{
//   			ResourceConfigurationId: jsii.String("resourceConfigurationId"),
//   		},
//   		ServiceManaged: &ServiceManagedModeProperty{
//   			HostAddress: jsii.String("hostAddress"),
//   			VpcId: jsii.String("vpcId"),
//
//   			// the properties below are optional
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
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Certificate: jsii.String("certificate"),
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
type CfnPrivateConnectionProps struct {
	// The connection configuration, either SelfManaged or ServiceManaged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html#cfn-devopsagent-privateconnection-connectionconfiguration
	//
	ConnectionConfiguration interface{} `field:"required" json:"connectionConfiguration" yaml:"connectionConfiguration"`
	// Unique name for this Private Connection within the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html#cfn-devopsagent-privateconnection-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Certificate for the Private Connection.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html#cfn-devopsagent-privateconnection-certificate
	//
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsagent-privateconnection.html#cfn-devopsagent-privateconnection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

