package awscdkapigatewayv2alpha


// Supported HTTP methods.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var bookStoreDefaultFn function
//
//
//   getBooksIntegration := awscdkapigatewayv2integrationsalpha.NewHttpUrlIntegration(jsii.String("GetBooksIntegration"), jsii.String("https://get-books-proxy.example.com"))
//   bookStoreDefaultIntegration := awscdkapigatewayv2integrationsalpha.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), bookStoreDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   	Integration: getBooksIntegration,
//   })
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []*httpMethod{
//   		apigwv2.*httpMethod_ANY,
//   	},
//   	Integration: bookStoreDefaultIntegration,
//   })
//
// Deprecated.
type HttpMethod string

const (
	// HTTP ANY.
	// Deprecated.
	HttpMethod_ANY HttpMethod = "ANY"
	// HTTP DELETE.
	// Deprecated.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// HTTP GET.
	// Deprecated.
	HttpMethod_GET HttpMethod = "GET"
	// HTTP HEAD.
	// Deprecated.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// HTTP OPTIONS.
	// Deprecated.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// HTTP PATCH.
	// Deprecated.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// HTTP POST.
	// Deprecated.
	HttpMethod_POST HttpMethod = "POST"
	// HTTP PUT.
	// Deprecated.
	HttpMethod_PUT HttpMethod = "PUT"
)

