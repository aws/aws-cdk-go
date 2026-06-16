package awsbedrockagentcore


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
//   	VpcIdentifier: jsii.String("vpcIdentifier"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html
//
type CfnGatewayTargetPropsMixin_ManagedVpcResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-endpointipaddresstype
	//
	EndpointIpAddressType *string `field:"optional" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-routingdomain
	//
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-vpcidentifier
	//
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}

