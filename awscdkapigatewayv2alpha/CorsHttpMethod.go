package awscdkapigatewayv2alpha


// Supported CORS HTTP methods.
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
type CorsHttpMethod string

const (
	// HTTP ANY.
	// Deprecated.
	CorsHttpMethod_ANY CorsHttpMethod = "ANY"
	// HTTP DELETE.
	// Deprecated.
	CorsHttpMethod_DELETE CorsHttpMethod = "DELETE"
	// HTTP GET.
	// Deprecated.
	CorsHttpMethod_GET CorsHttpMethod = "GET"
	// HTTP HEAD.
	// Deprecated.
	CorsHttpMethod_HEAD CorsHttpMethod = "HEAD"
	// HTTP OPTIONS.
	// Deprecated.
	CorsHttpMethod_OPTIONS CorsHttpMethod = "OPTIONS"
	// HTTP PATCH.
	// Deprecated.
	CorsHttpMethod_PATCH CorsHttpMethod = "PATCH"
	// HTTP POST.
	// Deprecated.
	CorsHttpMethod_POST CorsHttpMethod = "POST"
	// HTTP PUT.
	// Deprecated.
	CorsHttpMethod_PUT CorsHttpMethod = "PUT"
)

