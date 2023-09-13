package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

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
type IntegrationOptions struct {
	// A list of request parameters whose values are to be cached.
	//
	// It determines
	// request parameters that will make it into the cache key.
	CacheKeyParameters *[]*string `field:"optional" json:"cacheKeyParameters" yaml:"cacheKeyParameters"`
	// An API-specific tag group of related cached parameters.
	CacheNamespace *string `field:"optional" json:"cacheNamespace" yaml:"cacheNamespace"`
	// The type of network connection to the integration endpoint.
	// Default: - ConnectionType.VPC_LINK if `vpcLink` property is configured; ConnectionType.Internet otherwise.
	//
	ConnectionType ConnectionType `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Specifies how to handle request payload content type conversions.
	// Default: none if this property isn't defined, the request payload is passed
	// through from the method request to the integration request without
	// modification, provided that the `passthroughBehaviors` property is
	// configured to support payload pass-through.
	//
	ContentHandling ContentHandling `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Requires that the caller's identity be passed through from the request.
	// Default: Caller identity is not passed through.
	//
	CredentialsPassthrough *bool `field:"optional" json:"credentialsPassthrough" yaml:"credentialsPassthrough"`
	// An IAM role that API Gateway assumes.
	//
	// Mutually exclusive with `credentialsPassThrough`.
	// Default: A role is not assumed.
	//
	CredentialsRole awsiam.IRole `field:"optional" json:"credentialsRole" yaml:"credentialsRole"`
	// The response that API Gateway provides after a method's backend completes processing a request.
	//
	// API Gateway intercepts the response from the
	// backend so that you can control how API Gateway surfaces backend
	// responses. For example, you can map the backend status codes to codes
	// that you define.
	IntegrationResponses *[]*IntegrationResponse `field:"optional" json:"integrationResponses" yaml:"integrationResponses"`
	// Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource.
	//
	// There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and
	// NEVER.
	PassthroughBehavior PassthroughBehavior `field:"optional" json:"passthroughBehavior" yaml:"passthroughBehavior"`
	// The request parameters that API Gateway sends with the backend request.
	//
	// Specify request parameters as key-value pairs (string-to-string
	// mappings), with a destination as the key and a source as the value.
	//
	// Specify the destination by using the following pattern
	// integration.request.location.name, where location is querystring, path,
	// or header, and name is a valid, unique parameter name.
	//
	// The source must be an existing method request parameter or a static
	// value. You must enclose static values in single quotation marks and
	// pre-encode these values based on their destination in the request.
	RequestParameters *map[string]*string `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// A map of Apache Velocity templates that are applied on the request payload.
	//
	// The template that API Gateway uses is based on the value of the
	// Content-Type header that's sent by the client. The content type value is
	// the key, and the template is the value (specified as a string), such as
	// the following snippet:
	//
	// ```
	//   { "application/json": "{ \"statusCode\": 200 }" }
	// ```.
	// See: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
	//
	RequestTemplates *map[string]*string `field:"optional" json:"requestTemplates" yaml:"requestTemplates"`
	// The maximum amount of time an integration will run before it returns without a response.
	//
	// Must be between 50 milliseconds and 29 seconds.
	// Default: Duration.seconds(29)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The VpcLink used for the integration.
	//
	// Required if connectionType is VPC_LINK.
	VpcLink IVpcLink `field:"optional" json:"vpcLink" yaml:"vpcLink"`
}

