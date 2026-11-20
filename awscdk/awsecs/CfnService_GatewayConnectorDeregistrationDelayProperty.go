package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayConnectorDeregistrationDelayProperty := &GatewayConnectorDeregistrationDelayProperty{
//   	TimeoutSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorderegistrationdelay.html
//
type CfnService_GatewayConnectorDeregistrationDelayProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-gatewayconnectorderegistrationdelay.html#cfn-ecs-service-gatewayconnectorderegistrationdelay-timeoutseconds
	//
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

