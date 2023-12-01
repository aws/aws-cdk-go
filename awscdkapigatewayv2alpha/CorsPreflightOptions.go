package awscdkapigatewayv2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for the CORS Configuration.
//
// Example:
//   apigwv2.NewHttpApi(this, jsii.String("HttpProxyApi"), &HttpApiProps{
//   	CorsPreflight: &CorsPreflightOptions{
//   		AllowHeaders: []*string{
//   			jsii.String("Authorization"),
//   		},
//   		AllowMethods: []corsHttpMethod{
//   			apigwv2.*corsHttpMethod_GET,
//   			apigwv2.*corsHttpMethod_HEAD,
//   			apigwv2.*corsHttpMethod_OPTIONS,
//   			apigwv2.*corsHttpMethod_POST,
//   		},
//   		AllowOrigins: []*string{
//   			jsii.String("*"),
//   		},
//   		MaxAge: awscdk.Duration_Days(jsii.Number(10)),
//   	},
//   })
//
// Deprecated.
type CorsPreflightOptions struct {
	// Specifies whether credentials are included in the CORS request.
	// Default: false.
	//
	// Deprecated.
	AllowCredentials *bool `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Represents a collection of allowed headers.
	// Default: - No Headers are allowed.
	//
	// Deprecated.
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Represents a collection of allowed HTTP methods.
	// Default: - No Methods are allowed.
	//
	// Deprecated.
	AllowMethods *[]CorsHttpMethod `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Represents a collection of allowed origins.
	// Default: - No Origins are allowed.
	//
	// Deprecated.
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// Represents a collection of exposed headers.
	// Default: - No Expose Headers are allowed.
	//
	// Deprecated.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The duration that the browser should cache preflight request results.
	// Default: Duration.seconds(0)
	//
	// Deprecated.
	MaxAge awscdk.Duration `field:"optional" json:"maxAge" yaml:"maxAge"`
}

