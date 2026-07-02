package awsbedrockagentcore


// Configuration for a service-managed VPC endpoint.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html
//
type CfnHarness_ManagedVpcResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-endpointipaddresstype
	//
	EndpointIpAddressType *string `field:"required" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-vpcidentifier
	//
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-routingdomain
	//
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

