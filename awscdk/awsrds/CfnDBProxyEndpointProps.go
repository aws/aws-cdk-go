package awsrds


// Properties for defining a `CfnDBProxyEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDBProxyEndpointProps := &CfnDBProxyEndpointProps{
//   	DbProxyEndpointName: jsii.String("dbProxyEndpointName"),
//   	DbProxyName: jsii.String("dbProxyName"),
//   	VpcSubnetIds: []*string{
//   		jsii.String("vpcSubnetIds"),
//   	},
//
//   	// the properties below are optional
//   	EndpointNetworkType: jsii.String("endpointNetworkType"),
//   	Tags: []tagFormatProperty{
//   		&tagFormatProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TargetRole: jsii.String("targetRole"),
//   	VpcSecurityGroupIds: []*string{
//   		jsii.String("vpcSecurityGroupIds"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html
//
type CfnDBProxyEndpointProps struct {
	// The name of the DB proxy endpoint to create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-dbproxyendpointname
	//
	DbProxyEndpointName *string `field:"required" json:"dbProxyEndpointName" yaml:"dbProxyEndpointName"`
	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-dbproxyname
	//
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The VPC subnet IDs for the DB proxy endpoint that you create.
	//
	// You can specify a different set of subnet IDs than for the original DB proxy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-vpcsubnetids
	//
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// The network type of the DB proxy endpoint.
	//
	// The network type determines the IP version that the proxy endpoint supports.
	//
	// Valid values:
	//
	// - `IPV4` - The proxy endpoint supports IPv4 only.
	// - `IPV6` - The proxy endpoint supports IPv6 only.
	// - `DUAL` - The proxy endpoint supports both IPv4 and IPv6.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-endpointnetworktype
	//
	EndpointNetworkType *string `field:"optional" json:"endpointNetworkType" yaml:"endpointNetworkType"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-tags
	//
	Tags *[]*CfnDBProxyEndpoint_TagFormatProperty `field:"optional" json:"tags" yaml:"tags"`
	// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-targetrole
	//
	TargetRole *string `field:"optional" json:"targetRole" yaml:"targetRole"`
	// The VPC security group IDs for the DB proxy endpoint that you create.
	//
	// You can specify a different set of security group IDs than for the original DB proxy. The default is the default security group for the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxyendpoint.html#cfn-rds-dbproxyendpoint-vpcsecuritygroupids
	//
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

