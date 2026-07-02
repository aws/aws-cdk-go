package awsbedrockagentcore


// Private endpoint configuration for connecting to the OpenID Connect discovery endpoint over a private network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   privateEndpointProperty := &PrivateEndpointProperty{
//   	ManagedVpcResource: &ManagedVpcResourceProperty{
//   		EndpointIpAddressType: jsii.String("endpointIpAddressType"),
//   		RoutingDomain: jsii.String("routingDomain"),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   		Tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		VpcIdentifier: jsii.String("vpcIdentifier"),
//   	},
//   	SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   		ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-privateendpoint.html
//
type CfnHarnessPropsMixin_PrivateEndpointProperty struct {
	// Configuration for a service-managed VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-privateendpoint.html#cfn-bedrockagentcore-harness-privateendpoint-managedvpcresource
	//
	ManagedVpcResource interface{} `field:"optional" json:"managedVpcResource" yaml:"managedVpcResource"`
	// Configuration for connecting to a private resource using a self-managed VPC Lattice resource configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-privateendpoint.html#cfn-bedrockagentcore-harness-privateendpoint-selfmanagedlatticeresource
	//
	SelfManagedLatticeResource interface{} `field:"optional" json:"selfManagedLatticeResource" yaml:"selfManagedLatticeResource"`
}

