package previewawsbedrockagentcoremixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   interceptorConfigurationProperty := &InterceptorConfigurationProperty{
//   	Lambda: &LambdaInterceptorConfigurationProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorconfiguration.html
//
type CfnGatewayPropsMixin_InterceptorConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorconfiguration.html#cfn-bedrockagentcore-gateway-interceptorconfiguration-lambda
	//
	Lambda interface{} `field:"optional" json:"lambda" yaml:"lambda"`
}

