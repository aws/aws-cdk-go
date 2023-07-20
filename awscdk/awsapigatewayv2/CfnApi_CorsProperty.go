package awsapigatewayv2


// The `Cors` property specifies a CORS configuration for an API.
//
// Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsProperty := &CorsProperty{
//   	AllowCredentials: jsii.Boolean(false),
//   	AllowHeaders: []*string{
//   		jsii.String("allowHeaders"),
//   	},
//   	AllowMethods: []*string{
//   		jsii.String("allowMethods"),
//   	},
//   	AllowOrigins: []*string{
//   		jsii.String("allowOrigins"),
//   	},
//   	ExposeHeaders: []*string{
//   		jsii.String("exposeHeaders"),
//   	},
//   	MaxAge: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html
//
type CfnApi_CorsProperty struct {
	// Specifies whether credentials are included in the CORS request.
	//
	// Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-allowcredentials
	//
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Represents a collection of allowed headers.
	//
	// Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-allowheaders
	//
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Represents a collection of allowed HTTP methods.
	//
	// Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-allowmethods
	//
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Represents a collection of allowed origins.
	//
	// Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-alloworigins
	//
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// Represents a collection of exposed headers.
	//
	// Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-exposeheaders
	//
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The number of seconds that the browser should cache preflight request results.
	//
	// Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-maxage
	//
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

