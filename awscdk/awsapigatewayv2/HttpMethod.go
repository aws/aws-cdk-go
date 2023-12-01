package awsapigatewayv2


// Supported HTTP methods.
//
// Example:
//   import "github.com/aws-samples/dummy/awscdklib/awsapigatewayv2integrations"
//
//   var booksDefaultFn function
//
//
//   getBooksIntegration := awscdklibawsapigatewayv2integrations.NewHttpUrlIntegration(jsii.String("GetBooksIntegration"), jsii.String("https://get-books-proxy.example.com"))
//   booksDefaultIntegration := awscdklibawsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)
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
//   	Integration: booksDefaultIntegration,
//   })
//
type HttpMethod string

const (
	// HTTP ANY.
	HttpMethod_ANY HttpMethod = "ANY"
	// HTTP DELETE.
	HttpMethod_DELETE HttpMethod = "DELETE"
	// HTTP GET.
	HttpMethod_GET HttpMethod = "GET"
	// HTTP HEAD.
	HttpMethod_HEAD HttpMethod = "HEAD"
	// HTTP OPTIONS.
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
	// HTTP PATCH.
	HttpMethod_PATCH HttpMethod = "PATCH"
	// HTTP POST.
	HttpMethod_POST HttpMethod = "POST"
	// HTTP PUT.
	HttpMethod_PUT HttpMethod = "PUT"
)

