package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRouteServerEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRouteServerEndpointProps := &CfnRouteServerEndpointProps{
//   	RouteServerId: jsii.String("routeServerId"),
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverendpoint.html
//
type CfnRouteServerEndpointProps struct {
	// The ID of the route server associated with this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverendpoint.html#cfn-ec2-routeserverendpoint-routeserverid
	//
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// The ID of the subnet to place the route server endpoint into.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverendpoint.html#cfn-ec2-routeserverendpoint-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Any tags assigned to the route server endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-routeserverendpoint.html#cfn-ec2-routeserverendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

