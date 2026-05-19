package awsbedrockagentcore


// Configuration returned from binding an interceptor to a Gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var configuration interface{}
//
//   interceptorBindConfig := &InterceptorBindConfig{
//   	Configuration: configuration,
//   }
//
type InterceptorBindConfig struct {
	// The CloudFormation configuration for this interceptor.
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
}

