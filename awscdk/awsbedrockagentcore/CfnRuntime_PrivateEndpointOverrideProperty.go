package awsbedrockagentcore


// Override mapping of a domain to a private endpoint.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-privateendpointoverride.html
//
type CfnRuntime_PrivateEndpointOverrideProperty struct {
	// The domain to override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-privateendpointoverride.html#cfn-bedrockagentcore-runtime-privateendpointoverride-domain
	//
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Private endpoint configuration.
	//
	// Exactly one of SelfManagedLatticeResource or ManagedVpcResource must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-privateendpointoverride.html#cfn-bedrockagentcore-runtime-privateendpointoverride-privateendpoint
	//
	PrivateEndpoint interface{} `field:"required" json:"privateEndpoint" yaml:"privateEndpoint"`
}

