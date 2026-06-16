package awsbedrockagentcore


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
//   	},
//   	SelfManagedLatticeResource: &SelfManagedLatticeResourceProperty{
//   		ResourceConfigurationIdentifier: jsii.String("resourceConfigurationIdentifier"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-privateendpoint.html
//
type CfnGatewayTarget_PrivateEndpointProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-privateendpoint.html#cfn-bedrockagentcore-gatewaytarget-privateendpoint-managedvpcresource
	//
	ManagedVpcResource interface{} `field:"optional" json:"managedVpcResource" yaml:"managedVpcResource"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-privateendpoint.html#cfn-bedrockagentcore-gatewaytarget-privateendpoint-selfmanagedlatticeresource
	//
	SelfManagedLatticeResource interface{} `field:"optional" json:"selfManagedLatticeResource" yaml:"selfManagedLatticeResource"`
}

