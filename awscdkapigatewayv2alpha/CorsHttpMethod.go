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
// Experimental.
type CorsHttpMethod string

const (
	// HTTP ANY.
	// Experimental.
	CorsHttpMethod_ANY CorsHttpMethod = "ANY"
	// HTTP DELETE.
	// Experimental.
	CorsHttpMethod_DELETE CorsHttpMethod = "DELETE"
	// HTTP GET.
	// Experimental.
	CorsHttpMethod_GET CorsHttpMethod = "GET"
	// HTTP HEAD.
	// Experimental.
	CorsHttpMethod_HEAD CorsHttpMethod = "HEAD"
	// HTTP OPTIONS.
	// Experimental.
	CorsHttpMethod_OPTIONS CorsHttpMethod = "OPTIONS"
	// HTTP PATCH.
	// Experimental.
	CorsHttpMethod_PATCH CorsHttpMethod = "PATCH"
	// HTTP POST.
	// Experimental.
	CorsHttpMethod_POST CorsHttpMethod = "POST"
	// HTTP PUT.
	// Experimental.
	CorsHttpMethod_PUT CorsHttpMethod = "PUT"
)

