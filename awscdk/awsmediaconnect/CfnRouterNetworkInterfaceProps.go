package awsmediaconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRouterNetworkInterface`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouterNetworkInterfaceProps := &CfnRouterNetworkInterfaceProps{
//   	Configuration: &RouterNetworkInterfaceConfigurationProperty{
//   		Public: &PublicRouterNetworkInterfaceConfigurationProperty{
//   			AllowRules: []interface{}{
//   				&PublicRouterNetworkInterfaceRuleProperty{
//   					Cidr: jsii.String("cidr"),
//   				},
//   			},
//   		},
//   		Vpc: &VpcRouterNetworkInterfaceConfigurationProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	RegionName: jsii.String("regionName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routernetworkinterface.html
//
type CfnRouterNetworkInterfaceProps struct {
	// The configuration settings for a router network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routernetworkinterface.html#cfn-mediaconnect-routernetworkinterface-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// The name of the router network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routernetworkinterface.html#cfn-mediaconnect-routernetworkinterface-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS Region where the router network interface is located.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routernetworkinterface.html#cfn-mediaconnect-routernetworkinterface-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// Key-value pairs that can be used to tag and organize this router network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-routernetworkinterface.html#cfn-mediaconnect-routernetworkinterface-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

