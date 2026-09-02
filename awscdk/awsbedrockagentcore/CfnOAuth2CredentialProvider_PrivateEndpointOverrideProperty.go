package awsbedrockagentcore


// A mapping of a specific domain to a private endpoint for secure connectivity through a VPC Lattice resource configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateEndpointOverrideProperty := &PrivateEndpointOverrideProperty{
//   	Domain: jsii.String("domain"),
//   	PrivateEndpoint: &PrivateEndpointProperty{
//   		ManagedVpcResource: &ManagedVpcResourceProperty{
//   			EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   			VpcIdentifier: jsii.String("vpcIdentifier"),
//
//   			// the properties below are optional
//   			RoutingDomain: jsii.String("routingDomain"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			Tags: map[string]*string{
//   				"tagsKey": jsii.String("tags"),
//   			},
//   		},
//   		SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   			ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privateendpointoverride.html
//
type CfnOAuth2CredentialProvider_PrivateEndpointOverrideProperty struct {
	// The domain to override with a private endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privateendpointoverride.html#cfn-bedrockagentcore-oauth2credentialprovider-privateendpointoverride-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The private endpoint configuration for connecting to private resources in your VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-oauth2credentialprovider-privateendpointoverride.html#cfn-bedrockagentcore-oauth2credentialprovider-privateendpointoverride-privateendpoint
	//
	PrivateEndpoint interface{} `field:"required" json:"privateEndpoint" yaml:"privateEndpoint"`
}

