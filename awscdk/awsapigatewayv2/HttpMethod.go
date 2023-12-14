package awsapigatewayv2


// Supported HTTP methods.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var booksDefaultFn function
//
//   booksIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("BooksIntegration"), booksDefaultFn)
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []httpMethod{
//   		apigwv2.*httpMethod_GET,
//   	},
//   	Integration: booksIntegration,
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

