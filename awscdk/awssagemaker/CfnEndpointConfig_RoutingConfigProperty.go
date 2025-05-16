package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routingConfigProperty := &RoutingConfigProperty{
//   	RoutingStrategy: jsii.String("routingStrategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-routingconfig.html
//
type CfnEndpointConfig_RoutingConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpointconfig-routingconfig.html#cfn-sagemaker-endpointconfig-routingconfig-routingstrategy
	//
	RoutingStrategy *string `field:"optional" json:"routingStrategy" yaml:"routingStrategy"`
}

