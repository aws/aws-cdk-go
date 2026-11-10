package awsecs


// The header configuration for test traffic routing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayTestTrafficHeaderProperty := &GatewayTestTrafficHeaderProperty{
//   	Name: jsii.String("name"),
//   	Values: []*string{
//   		jsii.String("values"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewaytesttrafficheader.html
//
type CfnGateway_GatewayTestTrafficHeaderProperty struct {
	// The name of the header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewaytesttrafficheader.html#cfn-ecs-gateway-gatewaytesttrafficheader-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The values for the header.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-gateway-gatewaytesttrafficheader.html#cfn-ecs-gateway-gatewaytesttrafficheader-values
	//
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

