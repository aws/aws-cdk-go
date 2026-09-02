package awsbedrockagentcore


// The private endpoint configuration for connecting to private resources in your VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateEndpointProperty := &PrivateEndpointProperty{
//   	ManagedVpcResource: &ManagedVpcResourceProperty{
//   		EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   		// the properties below are optional
//   		RoutingDomain: jsii.String("routingDomain"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   	},
//   	SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   		ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privateendpoint.html
//
type CfnOAuth2CredentialProvider_PrivateEndpointProperty struct {
	// Configuration for a managed VPC Lattice resource.
	//
	// AgentCore creates and manages the VPC Lattice resource gateway and resource configuration on your behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privateendpoint.html#cfn-bedrockagentcore-oauth2credentialprovider-privateendpoint-managedvpcresource
	//
	ManagedVpcResource interface{} `field:"optional" json:"managedVpcResource" yaml:"managedVpcResource"`
	// Configuration for a self-managed VPC Lattice resource.
	//
	// You create and manage the VPC Lattice resource gateway and resource configuration, then provide the resource configuration identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privateendpoint.html#cfn-bedrockagentcore-oauth2credentialprovider-privateendpoint-selfmanagedlatticeresource
	//
	SelfManagedLatticeResource interface{} `field:"optional" json:"selfManagedLatticeResource" yaml:"selfManagedLatticeResource"`
}

