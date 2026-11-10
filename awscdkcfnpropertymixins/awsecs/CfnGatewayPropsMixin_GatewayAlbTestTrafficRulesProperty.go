package awsecs


// The test traffic rules for the gateway ALB.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   gatewayAlbTestTrafficRulesProperty := &GatewayAlbTestTrafficRulesProperty{
//   	Header: &GatewayTestTrafficHeaderProperty{
//   		Name: jsii.String("name"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayalbtesttrafficrules.html
//
type CfnGatewayPropsMixin_GatewayAlbTestTrafficRulesProperty struct {
	// The header configuration for test traffic routing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewayalbtesttrafficrules.html#cfn-ecs-gateway-gatewayalbtesttrafficrules-header
	//
	Header interface{} `field:"optional" json:"header" yaml:"header"`
}

