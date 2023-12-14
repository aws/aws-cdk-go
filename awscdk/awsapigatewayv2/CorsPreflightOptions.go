package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for the CORS Configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsPreflightOptions := &CorsPreflightOptions{
//   	AllowCredentials: jsii.Boolean(false),
//   	AllowHeaders: []*string{
//   		jsii.String("allowHeaders"),
//   	},
//   	AllowMethods: []corsHttpMethod{
//   		awscdk.Aws_apigatewayv2.*corsHttpMethod_ANY,
//   	},
//   	AllowOrigins: []*string{
//   		jsii.String("allowOrigins"),
//   	},
//   	ExposeHeaders: []*string{
//   		jsii.String("exposeHeaders"),
//   	},
//   	MaxAge: cdk.Duration_Minutes(jsii.Number(30)),
//   }
//
type CorsPreflightOptions struct {
	// Specifies whether credentials are included in the CORS request.
	// Default: false.
	//
	AllowCredentials *bool `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Represents a collection of allowed headers.
	// Default: - No Headers are allowed.
	//
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Represents a collection of allowed HTTP methods.
	// Default: - No Methods are allowed.
	//
	AllowMethods *[]CorsHttpMethod `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Represents a collection of allowed origins.
	// Default: - No Origins are allowed.
	//
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// Represents a collection of exposed headers.
	// Default: - No Expose Headers are allowed.
	//
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The duration that the browser should cache preflight request results.
	// Default: Duration.seconds(0)
	//
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
}

