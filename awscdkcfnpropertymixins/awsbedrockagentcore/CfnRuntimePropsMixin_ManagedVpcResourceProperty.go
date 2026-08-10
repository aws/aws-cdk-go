package awsbedrockagentcore


// Managed VPC resource configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   managedVpcResourceProperty := &ManagedVpcResourceProperty{
//   	EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   	RoutingDomain: jsii.String("routingDomain"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	VpcIdentifier: jsii.String("vpcIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html
//
type CfnRuntimePropsMixin_ManagedVpcResourceProperty struct {
	// The IP address type for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-endpointipaddresstype
	//
	EndpointIpAddressType *string `field:"optional" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// An intermediate domain to use as the resource configuration endpoint instead of the actual target domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-routingdomain
	//
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// The security group IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet IDs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Tags to apply to the managed VPC Lattice resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The VPC identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-managedvpcresource.html#cfn-bedrockagentcore-runtime-managedvpcresource-vpcidentifier
	//
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}

