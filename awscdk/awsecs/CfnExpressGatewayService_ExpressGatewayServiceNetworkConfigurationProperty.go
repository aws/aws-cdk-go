package awsecs


// The network configuration for an Express service.
//
// By default, an Express service utilizes subnets and security groups associated with the default VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   expressGatewayServiceNetworkConfigurationProperty := &ExpressGatewayServiceNetworkConfigurationProperty{
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	Subnets: []*string{
//   		jsii.String("subnets"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration.html
//
type CfnExpressGatewayService_ExpressGatewayServiceNetworkConfigurationProperty struct {
	// The IDs of the security groups associated with the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The IDs of the subnets associated with the Express service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration-subnets
	//
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

