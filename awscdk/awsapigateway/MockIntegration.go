package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// This type of integration lets API Gateway return a response without sending the request further to the backend.
//
// This is useful for API testing because it
// can be used to test the integration set up without incurring charges for
// using the backend and to enable collaborative development of an API. In
// collaborative development, a team can isolate their development effort by
// setting up simulations of API components owned by other teams by using the
// MOCK integrations. It is also used to return CORS-related headers to ensure
// that the API method permits CORS access. In fact, the API Gateway console
// integrates the OPTIONS method to support CORS with a mock integration.
// Gateway responses are other examples of mock integrations.
//
// Example:
//   import path "github.com/aws-samples/dummy/path"
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
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.AssetCode_FromAsset(path.join(__dirname, jsii.String("integ.request-authorizer.handler"))),
//   })
//
//   restapi := awscdk.NewRestApi(stack, jsii.String("MyRestApi"), &RestApiProps{
//   	CloudWatchRole: jsii.Boolean(true),
//   })
//
//   authorizer := awscdk.NewRequestAuthorizer(stack, jsii.String("MyAuthorizer"), &RequestAuthorizerProps{
//   	Handler: authorizerFn,
//   	IdentitySources: []*string{
//   		awscdk.IdentitySource_Header(jsii.String("Authorization")),
//   		awscdk.IdentitySource_QueryString(jsii.String("allow")),
//   	},
//   })
//
//   secondAuthorizer := awscdk.NewRequestAuthorizer(stack, jsii.String("MySecondAuthorizer"), &RequestAuthorizerProps{
//   	Handler: authorizerFn,
//   	IdentitySources: []*string{
//   		awscdk.IdentitySource_*Header(jsii.String("Authorization")),
//   		awscdk.IdentitySource_*QueryString(jsii.String("allow")),
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
//   restapi.Root.ResourceForPath(jsii.String("auth")).AddMethod(jsii.String("ANY"), awscdk.NewMockIntegration(&IntegrationOptions{
//   	IntegrationResponses: []*integrationResponse{
//   		&integrationResponse{
//   			StatusCode: jsii.String("200"),
//   		},
//   	},
//   	PassthroughBehavior: awscdk.PassthroughBehavior_NEVER,
//   	RequestTemplates: map[string]*string{
//   		"application/json": jsii.String("{ \"statusCode\": 200 }"),
//   	},
//   }), &MethodOptions{
//   	MethodResponses: []*methodResponse{
//   		&methodResponse{
//   			StatusCode: jsii.String("200"),
//   		},
//   	},
//   	Authorizer: secondAuthorizer,
//   })
//
type MockIntegration interface {
	Integration
	// Can be overridden by subclasses to allow the integration to interact with the method being integrated, access the REST API object, method ARNs, etc.
	Bind(_method Method) *IntegrationConfig
}

// The jsii proxy struct for MockIntegration
type jsiiProxy_MockIntegration struct {
	jsiiProxy_Integration
}

func NewMockIntegration(options *IntegrationOptions) MockIntegration {
	_init_.Initialize()

	if err := validateNewMockIntegrationParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_MockIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.MockIntegration",
		[]interface{}{options},
		&j,
	)

	return &j
}

func NewMockIntegration_Override(m MockIntegration, options *IntegrationOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigateway.MockIntegration",
		[]interface{}{options},
		m,
	)
}

func (m *jsiiProxy_MockIntegration) Bind(_method Method) *IntegrationConfig {
	if err := m.validateBindParameters(_method); err != nil {
		panic(err)
	}
	var returns *IntegrationConfig

	_jsii_.Invoke(
		m,
		"bind",
		[]interface{}{_method},
		&returns,
	)

	return returns
}

