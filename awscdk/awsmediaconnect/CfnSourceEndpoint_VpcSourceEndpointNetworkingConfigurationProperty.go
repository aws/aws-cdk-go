package awsmediaconnect


// The networking configuration for a source endpoint in a virtual private cloud (VPC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcSourceEndpointNetworkingConfigurationProperty := &VpcSourceEndpointNetworkingConfigurationProperty{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-vpcsourceendpointnetworkingconfiguration.html
//
type CfnSourceEndpoint_VpcSourceEndpointNetworkingConfigurationProperty struct {
	// The IDs of the security groups to associate with the source endpoint.
	//
	// You can specify from 1 to 5 security groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-vpcsourceendpointnetworkingconfiguration.html#cfn-mediaconnect-sourceendpoint-vpcsourceendpointnetworkingconfiguration-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet to use for the source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-vpcsourceendpointnetworkingconfiguration.html#cfn-mediaconnect-sourceendpoint-vpcsourceendpointnetworkingconfiguration-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

