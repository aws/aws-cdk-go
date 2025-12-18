package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaInterceptorConfigurationProperty := &LambdaInterceptorConfigurationProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-lambdainterceptorconfiguration.html
//
type CfnGateway_LambdaInterceptorConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-lambdainterceptorconfiguration.html#cfn-bedrockagentcore-gateway-lambdainterceptorconfiguration-arn
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}

