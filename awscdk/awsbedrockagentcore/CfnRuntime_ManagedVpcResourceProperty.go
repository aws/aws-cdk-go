package awsbedrockagentcore


// Managed VPC resource configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   managedVpcResourceProperty := &ManagedVpcResourceProperty{
//   	EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   	// the properties below are optional
//   	RoutingDomain: jsii.String("routingDomain"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html
//
type CfnRuntime_ManagedVpcResourceProperty struct {
	// The IP address type for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-endpointipaddresstype
	//
	EndpointIpAddressType *string `field:"required" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// The subnet IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The VPC identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-vpcidentifier
	//
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// An intermediate domain to use as the resource configuration endpoint instead of the actual target domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-routingdomain
	//
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// The security group IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Tags to apply to the managed VPC Lattice resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

