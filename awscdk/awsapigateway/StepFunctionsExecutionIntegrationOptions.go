package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Options when configuring Step Functions synchronous integration with Rest API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//   var vpcLink vpcLink
//
//   stepFunctionsExecutionIntegrationOptions := &StepFunctionsExecutionIntegrationOptions{
//   	Authorizer: jsii.Boolean(false),
//   	CacheKeyParameters: []*string{
//   		jsii.String("cacheKeyParameters"),
//   	},
//   	CacheNamespace: jsii.String("cacheNamespace"),
//   	ConnectionType: awscdk.Aws_apigateway.ConnectionType_INTERNET,
//   	ContentHandling: awscdk.*Aws_apigateway.ContentHandling_CONVERT_TO_BINARY,
//   	CredentialsPassthrough: jsii.Boolean(false),
//   	CredentialsRole: role,
//   	Headers: jsii.Boolean(false),
//   	IntegrationResponses: []integrationResponse{
//   		&integrationResponse{
//   			StatusCode: jsii.String("statusCode"),
//
//   			// the properties below are optional
//   			ContentHandling: awscdk.*Aws_apigateway.ContentHandling_CONVERT_TO_BINARY,
//   			ResponseParameters: map[string]*string{
//   				"responseParametersKey": jsii.String("responseParameters"),
//   			},
//   			ResponseTemplates: map[string]*string{
//   				"responseTemplatesKey": jsii.String("responseTemplates"),
//   			},
//   			SelectionPattern: jsii.String("selectionPattern"),
//   		},
//   	},
//   	PassthroughBehavior: awscdk.*Aws_apigateway.PassthroughBehavior_WHEN_NO_MATCH,
//   	Path: jsii.Boolean(false),
//   	Querystring: jsii.Boolean(false),
//   	RequestContext: &RequestContext{
//   		AccountId: jsii.Boolean(false),
//   		ApiId: jsii.Boolean(false),
//   		ApiKey: jsii.Boolean(false),
//   		AuthorizerPrincipalId: jsii.Boolean(false),
//   		Caller: jsii.Boolean(false),
//   		CognitoAuthenticationProvider: jsii.Boolean(false),
//   		CognitoAuthenticationType: jsii.Boolean(false),
//   		CognitoIdentityId: jsii.Boolean(false),
//   		CognitoIdentityPoolId: jsii.Boolean(false),
//   		HttpMethod: jsii.Boolean(false),
//   		RequestId: jsii.Boolean(false),
//   		ResourceId: jsii.Boolean(false),
//   		ResourcePath: jsii.Boolean(false),
//   		SourceIp: jsii.Boolean(false),
//   		Stage: jsii.Boolean(false),
//   		User: jsii.Boolean(false),
//   		UserAgent: jsii.Boolean(false),
//   		UserArn: jsii.Boolean(false),
//   	},
//   	RequestParameters: map[string]*string{
//   		"requestParametersKey": jsii.String("requestParameters"),
//   	},
//   	RequestTemplates: map[string]*string{
//   		"requestTemplatesKey": jsii.String("requestTemplates"),
//   	},
//   	Timeout: cdk.Duration_Minutes(jsii.Number(30)),
//   	UseDefaultMethodResponses: jsii.Boolean(false),
//   	VpcLink: vpcLink,
//   }
//
type StepFunctionsExecutionIntegrationOptions struct {
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
	// If the whole authorizer object, including custom context values should be in the execution input.
	//
	// The execution input will include a new key `authorizer`:
	//
	// {
	//   "body": {},
	//   "authorizer": {
	//     "key": "value"
	//   }
	// }.
	// Default: false.
	//
	Authorizer *bool `field:"optional" json:"authorizer" yaml:"authorizer"`
	// Check if header is to be included inside the execution input.
	//
	// The execution input will include a new key `headers`:
	//
	// {
	//   "body": {},
	//   "headers": {
	//      "header1": "value",
	//      "header2": "value"
	//   }
	// }.
	// Default: false.
	//
	Headers *bool `field:"optional" json:"headers" yaml:"headers"`
	// Check if path is to be included inside the execution input.
	//
	// The execution input will include a new key `path`:
	//
	// {
	//   "body": {},
	//   "path": {
	//     "resourceName": "resourceValue"
	//   }
	// }.
	// Default: true.
	//
	Path *bool `field:"optional" json:"path" yaml:"path"`
	// Check if querystring is to be included inside the execution input.
	//
	// The execution input will include a new key `queryString`:
	//
	// {
	//   "body": {},
	//   "querystring": {
	//     "key": "value"
	//   }
	// }.
	// Default: true.
	//
	Querystring *bool `field:"optional" json:"querystring" yaml:"querystring"`
	// Which details of the incoming request must be passed onto the underlying state machine, such as, account id, user identity, request id, etc.
	//
	// The execution input will include a new key `requestContext`:
	//
	// {
	//   "body": {},
	//   "requestContext": {
	//       "key": "value"
	//   }
	// }.
	// Default: - all parameters within request context will be set as false.
	//
	RequestContext *RequestContext `field:"optional" json:"requestContext" yaml:"requestContext"`
	// Whether to add default response models with 200, 400, and 500 status codes to the method.
	// Default: true.
	//
	UseDefaultMethodResponses *bool `field:"optional" json:"useDefaultMethodResponses" yaml:"useDefaultMethodResponses"`
}

