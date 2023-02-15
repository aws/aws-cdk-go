package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specifies a cross-origin access property for a function URL.
//
// Example:
//   var fn function
//
//
//   fn.addFunctionUrl(&functionUrlOptions{
//   	authType: lambda.functionUrlAuthType_NONE,
//   	cors: &functionUrlCorsOptions{
//   		// Allow this to be called from websites on https://example.com.
//   		// Can also be ['*'] to allow all domain.
//   		allowedOrigins: []*string{
//   			jsii.String("https://example.com"),
//   		},
//   	},
//   })
//
type FunctionUrlCorsOptions struct {
	// Whether to allow cookies or other credentials in requests to your function URL.
	AllowCredentials *bool `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Headers that are specified in the Access-Control-Request-Headers header.
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// An HTTP method that you allow the origin to execute.
	AllowedMethods *[]HttpMethod `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more origins you want customers to be able to access the bucket from.
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// One or more headers in the response that you want customers to be able to access from their applications.
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
}

