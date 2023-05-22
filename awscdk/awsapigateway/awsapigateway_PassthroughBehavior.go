package awsapigateway


// Example:
//   import path "github.com/aws-samples/dummy/path"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // Against the RestApi endpoint from the stack output, run
//   // `curl -s -o /dev/null -w "%{http_code}" <url>` should return 401
//   // `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: deny' <url>?allow=yes` should return 403
//   // `curl -s -o /dev/null -w "%{http_code}" -H 'Authorization: allow' <url>?allow=yes` should return 200
//
//   app := awscdk.NewApp()
//   stack := awscdk.NewStack(app, jsii.String("RequestAuthorizerInteg"))
//
//   authorizerFn := lambda.NewFunction(stack, jsii.String("MyAuthorizerFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.AssetCode_FromAsset(path.join(__dirname, jsii.String("integ.request-authorizer.handler"))),
//   })
//
//   restapi := awscdk.NewRestApi(stack, jsii.String("MyRestApi"))
//
//   authorizer := awscdk.NewRequestAuthorizer(stack, jsii.String("MyAuthorizer"), &RequestAuthorizerProps{
//   	Handler: authorizerFn,
//   	IdentitySources: []*string{
//   		awscdk.IdentitySource_Header(jsii.String("Authorization")),
//   		awscdk.IdentitySource_QueryString(jsii.String("allow")),
//   	},
//   })
//
//   restapi.Root.AddMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&IntegrationOptions{
//   	IntegrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			StatusCode: jsii.String("200"),
//   		},
//   	},
//   	PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   	RequestTemplates: map[string]*string{
//   		"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   	},
//   }), &MethodOptions{
//   	MethodResponses: []methodResponse{
//   		&methodResponse{
//   			StatusCode: jsii.String("200"),
//   		},
//   	},
//   	Authorizer: Authorizer,
//   })
//
// Experimental.
type PassthroughBehavior string

const (
	// Passes the request body for unmapped content types through to the integration back end without transformation.
	// Experimental.
	PassthroughBehavior_WHEN_NO_MATCH PassthroughBehavior = "WHEN_NO_MATCH"
	// Rejects unmapped content types with an HTTP 415 'Unsupported Media Type' response.
	// Experimental.
	PassthroughBehavior_NEVER PassthroughBehavior = "NEVER"
	// Allows pass-through when the integration has NO content types mapped to templates.
	//
	// However if there is at least one content type defined,
	// unmapped content types will be rejected with the same 415 response.
	// Experimental.
	PassthroughBehavior_WHEN_NO_TEMPLATES PassthroughBehavior = "WHEN_NO_TEMPLATES"
)

