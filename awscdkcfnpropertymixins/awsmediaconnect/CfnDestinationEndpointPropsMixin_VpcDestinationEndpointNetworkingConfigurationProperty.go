package awsmediaconnect


// The networking configuration for a destination endpoint in a virtual private cloud (VPC).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   vpcDestinationEndpointNetworkingConfigurationProperty := &VpcDestinationEndpointNetworkingConfigurationProperty{
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-vpcdestinationendpointnetworkingconfiguration.html
//
type CfnDestinationEndpointPropsMixin_VpcDestinationEndpointNetworkingConfigurationProperty struct {
	// The IDs of the security groups to associate with the destination endpoint.
	//
	// You can specify from 1 to 5 security groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-vpcdestinationendpointnetworkingconfiguration.html#cfn-mediaconnect-destinationendpoint-vpcdestinationendpointnetworkingconfiguration-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet to use for the destination endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-vpcdestinationendpointnetworkingconfiguration.html#cfn-mediaconnect-destinationendpoint-vpcdestinationendpointnetworkingconfiguration-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

