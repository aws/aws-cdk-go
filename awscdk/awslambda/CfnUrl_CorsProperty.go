package awslambda


// The [Cross-Origin Resource Sharing (CORS)](https://docs.aws.amazon.com/https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) settings for your function URL. Use CORS to grant access to your function URL from any origin. You can also use CORS to control access for specific HTTP headers and methods in requests to your function URL.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-url-cors.html
//
type CfnUrl_CorsProperty struct {
	// Whether you want to allow cookies or other credentials in requests to your function URL.
	//
	// The default is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-url-cors.html#cfn-lambda-url-cors-allowcredentials
	//
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// The HTTP headers that origins can include in requests to your function URL.
	//
	// For example: `Date` , `Keep-Alive` , `X-Custom-Header` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-url-cors.html#cfn-lambda-url-cors-allowheaders
	//
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// The HTTP methods that are allowed when calling your function URL.
	//
	// For example: `GET` , `POST` , `DELETE` , or the wildcard character ( `*` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-url-cors.html#cfn-lambda-url-cors-allowmethods
	//
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// The origins that can access your function URL.
	//
	// You can list any number of specific origins, separated by a comma. For example: `https://www.example.com` , `http://localhost:60905` .
	//
	// Alternatively, you can grant access to all origins with the wildcard character ( `*` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-url-cors.html#cfn-lambda-url-cors-alloworigins
	//
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// The HTTP headers in your function response that you want to expose to origins that call your function URL.
	//
	// For example: `Date` , `Keep-Alive` , `X-Custom-Header` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-url-cors.html#cfn-lambda-url-cors-exposeheaders
	//
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// The maximum amount of time, in seconds, that browsers can cache results of a preflight request.
	//
	// By default, this is set to `0` , which means the browser will not cache results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-url-cors.html#cfn-lambda-url-cors-maxage
	//
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

