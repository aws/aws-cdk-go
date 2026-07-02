package awsbedrockagentcore


// Maps a domain to a private endpoint for resolving that domain over a private network.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-privateendpointoverride.html
//
type CfnHarness_PrivateEndpointOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-privateendpointoverride.html#cfn-bedrockagentcore-harness-privateendpointoverride-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Private endpoint configuration for connecting to the OpenID Connect discovery endpoint over a private network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-privateendpointoverride.html#cfn-bedrockagentcore-harness-privateendpointoverride-privateendpoint
	//
	PrivateEndpoint interface{} `field:"required" json:"privateEndpoint" yaml:"privateEndpoint"`
}

