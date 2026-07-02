package awsbedrockagentcore


// Configuration for a service-managed VPC endpoint.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html
//
type CfnHarnessPropsMixin_ManagedVpcResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-endpointipaddresstype
	//
	EndpointIpAddressType *string `field:"optional" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-routingdomain
	//
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-managedvpcresource.html#cfn-bedrockagentcore-harness-managedvpcresource-vpcidentifier
	//
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}

