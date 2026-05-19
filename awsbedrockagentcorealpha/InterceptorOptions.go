package awsbedrockagentcorealpha


// Options for configuring an interceptor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   interceptorOptions := &InterceptorOptions{
//   	PassRequestHeaders: jsii.Boolean(false),
//   }
//
// Experimental.
type InterceptorOptions struct {
	// Whether to pass request headers to the interceptor Lambda function.
	//
	// **Security Warning**: Request headers can contain sensitive information such as
	// authentication tokens and credentials. Only enable this if your interceptor needs
	// access to headers and you have verified that sensitive information is not logged
	// or exposed.
	// Default: false - Headers are not passed to interceptor for security.
	//
	// Experimental.
	PassRequestHeaders *bool `field:"optional" json:"passRequestHeaders" yaml:"passRequestHeaders"`
}

