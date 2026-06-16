package awsbedrockagentcore


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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html
//
type CfnGatewayTarget_ManagedVpcResourceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-endpointipaddresstype
	//
	EndpointIpAddressType *string `field:"required" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-vpcidentifier
	//
	VpcIdentifier *string `field:"required" json:"vpcIdentifier" yaml:"vpcIdentifier"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-routingdomain
	//
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gatewaytarget-managedvpcresource.html#cfn-bedrockagentcore-gatewaytarget-managedvpcresource-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
}

