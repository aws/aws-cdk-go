package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options for the CORS Configuration.
//
// Example:
//   apigwv2.NewHttpApi(this, jsii.String("HttpProxyApi"), &httpApiProps{
//   	corsPreflight: &corsPreflightOptions{
//   		allowHeaders: []*string{
//   			jsii.String("Authorization"),
//   		},
//   		allowMethods: []corsHttpMethod{
//   			apigwv2.*corsHttpMethod_GET,
//   			apigwv2.*corsHttpMethod_HEAD,
//   			apigwv2.*corsHttpMethod_OPTIONS,
//   			apigwv2.*corsHttpMethod_POST,
//   		},
//   		allowOrigins: []*string{
//   			jsii.String("*"),
//   		},
//   		maxAge: awscdk.Duration.days(jsii.Number(10)),
//   	},
//   })
//
// Experimental.
type CorsPreflightOptions struct {
	// Specifies whether credentials are included in the CORS request.
	// Experimental.
	AllowCredentials *bool `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Represents a collection of allowed headers.
	// Experimental.
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Represents a collection of allowed HTTP methods.
	// Experimental.
	AllowMethods *[]CorsHttpMethod `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Represents a collection of allowed origins.
	// Experimental.
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// Represents a collection of exposed headers.
	// Experimental.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The duration that the browser should cache preflight request results.
	// Experimental.
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
}

