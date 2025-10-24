package awsapigatewayv2


// Supported CORS HTTP methods.
//
// Example:
//   apigwv2.NewHttpApi(this, jsii.String("HttpProxyApi"), &HttpApiProps{
//   	CorsPreflight: &CorsPreflightOptions{
//   		AllowHeaders: []*string{
//   			jsii.String("Authorization"),
//   		},
//   		AllowMethods: []CorsHttpMethod{
//   			apigwv2.CorsHttpMethod_GET,
//   			apigwv2.CorsHttpMethod_HEAD,
//   			apigwv2.CorsHttpMethod_OPTIONS,
//   			apigwv2.CorsHttpMethod_POST,
//   		},
//   		AllowOrigins: []*string{
//   			jsii.String("*"),
//   		},
//   		MaxAge: awscdk.Duration_Days(jsii.Number(10)),
//   	},
//   })
//
type CorsHttpMethod string

const (
	// HTTP ANY.
	CorsHttpMethod_ANY CorsHttpMethod = "ANY"
	// HTTP DELETE.
	CorsHttpMethod_DELETE CorsHttpMethod = "DELETE"
	// HTTP GET.
	CorsHttpMethod_GET CorsHttpMethod = "GET"
	// HTTP HEAD.
	CorsHttpMethod_HEAD CorsHttpMethod = "HEAD"
	// HTTP OPTIONS.
	CorsHttpMethod_OPTIONS CorsHttpMethod = "OPTIONS"
	// HTTP PATCH.
	CorsHttpMethod_PATCH CorsHttpMethod = "PATCH"
	// HTTP POST.
	CorsHttpMethod_POST CorsHttpMethod = "POST"
	// HTTP PUT.
	CorsHttpMethod_PUT CorsHttpMethod = "PUT"
)

