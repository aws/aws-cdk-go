package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnExpressGatewayServicePropsMixin_ExpressGatewayServiceNetworkConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration.html#cfn-ecs-expressgatewayservice-expressgatewayservicenetworkconfiguration-subnets
	//
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

