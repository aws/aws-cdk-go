package awsbedrockagentcorealpha


// Configuration returned from binding an interceptor to a Gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   var configuration interface{}
//
//   interceptorBindConfig := &InterceptorBindConfig{
//   	Configuration: configuration,
//   }
//
// Experimental.
type InterceptorBindConfig struct {
	// The CloudFormation configuration for this interceptor.
	// Experimental.
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
}

