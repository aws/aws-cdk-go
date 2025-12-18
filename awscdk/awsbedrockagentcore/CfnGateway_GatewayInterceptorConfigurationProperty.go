package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayInterceptorConfigurationProperty := &GatewayInterceptorConfigurationProperty{
//   	InterceptionPoints: []*string{
//   		jsii.String("interceptionPoints"),
//   	},
//   	Interceptor: &InterceptorConfigurationProperty{
//   		Lambda: &LambdaInterceptorConfigurationProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//
//   	// the properties below are optional
//   	InputConfiguration: &InterceptorInputConfigurationProperty{
//   		PassRequestHeaders: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.html
//
type CfnGateway_GatewayInterceptorConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.html#cfn-bedrockagentcore-gateway-gatewayinterceptorconfiguration-interceptionpoints
	//
	InterceptionPoints *[]*string `field:"required" json:"interceptionPoints" yaml:"interceptionPoints"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.html#cfn-bedrockagentcore-gateway-gatewayinterceptorconfiguration-interceptor
	//
	Interceptor interface{} `field:"required" json:"interceptor" yaml:"interceptor"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.html#cfn-bedrockagentcore-gateway-gatewayinterceptorconfiguration-inputconfiguration
	//
	InputConfiguration interface{} `field:"optional" json:"inputConfiguration" yaml:"inputConfiguration"`
}

