package previewawsmedialivemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnNetworkPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNetworkMixinProps := &CfnNetworkMixinProps{
//   	IpPools: []interface{}{
//   		&IpPoolProperty{
//   			Cidr: jsii.String("cidr"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Routes: []interface{}{
//   		&RouteProperty{
//   			Cidr: jsii.String("cidr"),
//   			Gateway: jsii.String("gateway"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-network.html
//
type CfnNetworkMixinProps struct {
	// The list of IP address cidr pools for the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-network.html#cfn-medialive-network-ippools
	//
	IpPools interface{} `field:"optional" json:"ipPools" yaml:"ipPools"`
	// The user-specified name of the Network to be created.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-network.html#cfn-medialive-network-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The routes for the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-network.html#cfn-medialive-network-routes
	//
	Routes interface{} `field:"optional" json:"routes" yaml:"routes"`
	// A collection of key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-medialive-network.html#cfn-medialive-network-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

