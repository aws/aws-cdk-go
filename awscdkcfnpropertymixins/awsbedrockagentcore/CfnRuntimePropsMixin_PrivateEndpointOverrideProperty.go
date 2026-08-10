package awsbedrockagentcore


// Override mapping of a domain to a private endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   privateEndpointOverrideProperty := &PrivateEndpointOverrideProperty{
//   	Domain: jsii.String("domain"),
//   	PrivateEndpoint: &PrivateEndpointProperty{
//   		ManagedVpcResource: &ManagedVpcResourceProperty{
//   			EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   			RoutingDomain: jsii.String("routingDomain"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   			Tags: map[string]*string{
//   				"tagsKey": jsii.String("tags"),
//   			},
//   			VpcIdentifier: jsii.String("vpcIdentifier"),
//   		},
//   		SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   			ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-privateendpointoverride.html
//
type CfnRuntimePropsMixin_PrivateEndpointOverrideProperty struct {
	// The domain to override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-privateendpointoverride.html#cfn-bedrockagentcore-runtime-privateendpointoverride-domain
	//
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Private endpoint configuration.
	//
	// Exactly one of SelfManagedLatticeResource or ManagedVpcResource must be specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-runtime-privateendpointoverride.html#cfn-bedrockagentcore-runtime-privateendpointoverride-privateendpoint
	//
	PrivateEndpoint interface{} `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
}

