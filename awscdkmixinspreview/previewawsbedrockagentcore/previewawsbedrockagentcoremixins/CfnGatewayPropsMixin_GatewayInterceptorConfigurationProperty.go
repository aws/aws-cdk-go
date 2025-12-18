package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gatewayInterceptorConfigurationProperty := &GatewayInterceptorConfigurationProperty{
//   	InputConfiguration: &InterceptorInputConfigurationProperty{
//   		PassRequestHeaders: jsii.Boolean(false),
//   	},
//   	InterceptionPoints: []*string{
//   		jsii.String("interceptionPoints"),
//   	},
//   	Interceptor: &InterceptorConfigurationProperty{
//   		Lambda: &LambdaInterceptorConfigurationProperty{
//   			Arn: jsii.String("arn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.html
//
type CfnGatewayPropsMixin_GatewayInterceptorConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.html#cfn-bedrockagentcore-gateway-gatewayinterceptorconfiguration-inputconfiguration
	//
	InputConfiguration interface{} `field:"optional" json:"inputConfiguration" yaml:"inputConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.html#cfn-bedrockagentcore-gateway-gatewayinterceptorconfiguration-interceptionpoints
	//
	InterceptionPoints *[]*string `field:"optional" json:"interceptionPoints" yaml:"interceptionPoints"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-gatewayinterceptorconfiguration.html#cfn-bedrockagentcore-gateway-gatewayinterceptorconfiguration-interceptor
	//
	Interceptor interface{} `field:"optional" json:"interceptor" yaml:"interceptor"`
}

