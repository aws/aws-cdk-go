package awsapigateway


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   corsProperty := &corsProperty{
//   	allowCredentials: jsii.Boolean(false),
//   	allowHeaders: []*string{
//   		jsii.String("allowHeaders"),
//   	},
//   	allowMethods: []*string{
//   		jsii.String("allowMethods"),
//   	},
//   	allowOrigins: []*string{
//   		jsii.String("allowOrigins"),
//   	},
//   	exposeHeaders: []*string{
//   		jsii.String("exposeHeaders"),
//   	},
//   	maxAge: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnApiV2_CorsProperty struct {
	// `CfnApiV2.CorsProperty.AllowCredentials`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-allowcredentials
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// `CfnApiV2.CorsProperty.AllowHeaders`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-allowheaders
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// `CfnApiV2.CorsProperty.AllowMethods`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-allowmethods
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// `CfnApiV2.CorsProperty.AllowOrigins`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-alloworigins
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// `CfnApiV2.CorsProperty.ExposeHeaders`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-exposeheaders
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
	// `CfnApiV2.CorsProperty.MaxAge`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigatewayv2-api-cors.html#cfn-apigatewayv2-api-cors-maxage
	//
	// Deprecated: moved to package aws-apigatewayv2.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
}

