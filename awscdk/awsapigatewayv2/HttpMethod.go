package awsapigatewayv2


// Supported HTTP methods.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var booksDefaultFn Function
//
//
//   httpApi := apigwv2.NewHttpApi(this, jsii.String("HttpApi"))
//
//   getBooksIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("GetBooksIntegration"), booksDefaultFn, &HttpLambdaIntegrationProps{
//   	ScopePermissionToRoute: jsii.Boolean(false),
//   })
//   createBookIntegration := awscdk.NewHttpLambdaIntegration(jsii.String("CreateBookIntegration"), booksDefaultFn, &HttpLambdaIntegrationProps{
//   	ScopePermissionToRoute: jsii.Boolean(false),
//   })
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_GET,
//   	},
//   	Integration: getBooksIntegration,
//   })
//
//   httpApi.AddRoutes(&AddRoutesOptions{
//   	Path: jsii.String("/books"),
//   	Methods: []HttpMethod{
//   		apigwv2.HttpMethod_POST,
//   	},
//   	Integration: createBookIntegration,
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

