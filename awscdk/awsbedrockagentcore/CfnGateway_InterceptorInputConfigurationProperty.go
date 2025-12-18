package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   interceptorInputConfigurationProperty := &InterceptorInputConfigurationProperty{
//   	PassRequestHeaders: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorinputconfiguration.html
//
type CfnGateway_InterceptorInputConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-gateway-interceptorinputconfiguration.html#cfn-bedrockagentcore-gateway-interceptorinputconfiguration-passrequestheaders
	//
	PassRequestHeaders interface{} `field:"required" json:"passRequestHeaders" yaml:"passRequestHeaders"`
}

