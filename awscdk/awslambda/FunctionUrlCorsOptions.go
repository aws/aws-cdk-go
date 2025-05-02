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
//   fn.AddFunctionUrl(&FunctionUrlOptions{
//   	AuthType: lambda.FunctionUrlAuthType_NONE,
//   	Cors: &FunctionUrlCorsOptions{
//   		// Allow this to be called from websites on https://example.com.
//   		// Can also be ['*'] to allow all domain.
//   		AllowedOrigins: []*string{
//   			jsii.String("https://example.com"),
//   		},
//   	},
//   })
//
type FunctionUrlCorsOptions struct {
	// Whether to allow cookies or other credentials in requests to your function URL.
	// Default: false.
	//
	AllowCredentials *bool `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Headers that are specified in the Access-Control-Request-Headers header.
	// Default: - No headers allowed.
	//
	AllowedHeaders *[]*string `field:"optional" json:"allowedHeaders" yaml:"allowedHeaders"`
	// An HTTP method that you allow the origin to execute.
	// Default: - [HttpMethod.ALL]
	//
	AllowedMethods *[]HttpMethod `field:"optional" json:"allowedMethods" yaml:"allowedMethods"`
	// One or more origins you want customers to be able to access the bucket from.
	// Default: - No origins allowed.
	//
	AllowedOrigins *[]*string `field:"optional" json:"allowedOrigins" yaml:"allowedOrigins"`
	// One or more headers in the response that you want customers to be able to access from their applications.
	// Default: - No headers exposed.
	//
	ExposedHeaders *[]*string `field:"optional" json:"exposedHeaders" yaml:"exposedHeaders"`
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	// Default: - Browser default of 5 seconds.
	//
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
}

