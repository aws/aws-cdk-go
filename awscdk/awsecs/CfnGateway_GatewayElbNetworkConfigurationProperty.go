package awsecs


// The network configuration for the gateway ALB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayElbNetworkConfigurationProperty := &GatewayElbNetworkConfigurationProperty{
//   	InboundSecurityGroup: jsii.String("inboundSecurityGroup"),
//   	IpAddressType: jsii.String("ipAddressType"),
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbnetworkconfiguration.html
//
type CfnGateway_GatewayElbNetworkConfigurationProperty struct {
	// The inbound security group for the ALB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbnetworkconfiguration.html#cfn-ecs-gateway-gatewayelbnetworkconfiguration-inboundsecuritygroup
	//
	InboundSecurityGroup *string `field:"optional" json:"inboundSecurityGroup" yaml:"inboundSecurityGroup"`
	// The IP address type for the ALB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbnetworkconfiguration.html#cfn-ecs-gateway-gatewayelbnetworkconfiguration-ipaddresstype
	//
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// The subnets for the ALB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayelbnetworkconfiguration.html#cfn-ecs-gateway-gatewayelbnetworkconfiguration-subnets
	//
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

