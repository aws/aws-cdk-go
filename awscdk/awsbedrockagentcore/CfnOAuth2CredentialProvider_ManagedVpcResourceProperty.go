package awsbedrockagentcore


// Configuration for a managed VPC Lattice resource.
//
// AgentCore creates and manages the VPC Lattice resource gateway and resource configuration on your behalf.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-managedvpcresource.html
//
type CfnOAuth2CredentialProvider_ManagedVpcResourceProperty struct {
	// The IP address type for the resource configuration endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-managedvpcresource.html#cfn-bedrockagentcore-oauth2credentialprovider-managedvpcresource-endpointipaddresstype
	//
	EndpointIpAddressType *string `field:"required" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// The subnet IDs within the VPC where the VPC Lattice resource gateway is placed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-managedvpcresource.html#cfn-bedrockagentcore-oauth2credentialprovider-managedvpcresource-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the VPC that contains your private resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-managedvpcresource.html#cfn-bedrockagentcore-oauth2credentialprovider-managedvpcresource-vpcidentifier
	//
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// An intermediate publicly resolvable domain used as the VPC Lattice resource configuration endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-managedvpcresource.html#cfn-bedrockagentcore-oauth2credentialprovider-managedvpcresource-routingdomain
	//
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// The security group IDs to associate with the VPC Lattice resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-managedvpcresource.html#cfn-bedrockagentcore-oauth2credentialprovider-managedvpcresource-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// A map of tags (key-value pairs) to apply to a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-managedvpcresource.html#cfn-bedrockagentcore-oauth2credentialprovider-managedvpcresource-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

