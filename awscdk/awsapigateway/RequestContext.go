package awsapigateway


// Configure what must be included in the `requestContext`.
//
// More details can be found at mapping templates documentation.
//
// Example:
//   apigateway.NewStepFunctionsRestApi(this, jsii.String("StepFunctionsRestApi"), &StepFunctionsRestApiProps{
//   	StateMachine: machine,
//   	Headers: jsii.Boolean(true),
//   	Path: jsii.Boolean(false),
//   	Querystring: jsii.Boolean(false),
//   	Authorizer: jsii.Boolean(false),
//   	RequestContext: &RequestContext{
//   		Caller: jsii.Boolean(true),
//   		User: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html
//
type RequestContext struct {
	// Represents the information of $context.identity.accountId.
	//
	// Whether the AWS account of the API owner should be included in the request context.
	// Default: false.
	//
	AccountId *bool `field:"optional" json:"accountId" yaml:"accountId"`
	// Represents the information of $context.apiId.
	//
	// Whether the identifier API Gateway assigns to your API should be included in the request context.
	// Default: false.
	//
	ApiId *bool `field:"optional" json:"apiId" yaml:"apiId"`
	// Represents the information of $context.identity.apiKey.
	//
	// Whether the API key associated with the request should be included in request context.
	// Default: false.
	//
	ApiKey *bool `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Represents the information of $context.authorizer.principalId.
	//
	// Whether the principal user identifier associated with the token sent by the client and returned
	// from an API Gateway Lambda authorizer should be included in the request context.
	// Default: false.
	//
	AuthorizerPrincipalId *bool `field:"optional" json:"authorizerPrincipalId" yaml:"authorizerPrincipalId"`
	// Represents the information of $context.identity.caller.
	//
	// Whether the principal identifier of the caller that signed the request should be included in the request context.
	// Supported for resources that use IAM authorization.
	// Default: false.
	//
	Caller *bool `field:"optional" json:"caller" yaml:"caller"`
	// Represents the information of $context.identity.cognitoAuthenticationProvider.
	//
	// Whether the list of the Amazon Cognito authentication providers used by the caller making the request should be included in the request context.
	// Available only if the request was signed with Amazon Cognito credentials.
	// Default: false.
	//
	CognitoAuthenticationProvider *bool `field:"optional" json:"cognitoAuthenticationProvider" yaml:"cognitoAuthenticationProvider"`
	// Represents the information of $context.identity.cognitoAuthenticationType.
	//
	// Whether the Amazon Cognito authentication type of the caller making the request should be included in the request context.
	// Available only if the request was signed with Amazon Cognito credentials.
	// Possible values include authenticated for authenticated identities and unauthenticated for unauthenticated identities.
	// Default: false.
	//
	CognitoAuthenticationType *bool `field:"optional" json:"cognitoAuthenticationType" yaml:"cognitoAuthenticationType"`
	// Represents the information of $context.identity.cognitoIdentityId.
	//
	// Whether the Amazon Cognito identity ID of the caller making the request should be included in the request context.
	// Available only if the request was signed with Amazon Cognito credentials.
	// Default: false.
	//
	CognitoIdentityId *bool `field:"optional" json:"cognitoIdentityId" yaml:"cognitoIdentityId"`
	// Represents the information of $context.identity.cognitoIdentityPoolId.
	//
	// Whether the Amazon Cognito identity pool ID of the caller making the request should be included in the request context.
	// Available only if the request was signed with Amazon Cognito credentials.
	// Default: false.
	//
	CognitoIdentityPoolId *bool `field:"optional" json:"cognitoIdentityPoolId" yaml:"cognitoIdentityPoolId"`
	// Represents the information of $context.httpMethod.
	//
	// Whether the HTTP method used should be included in the request context.
	// Valid values include: DELETE, GET, HEAD, OPTIONS, PATCH, POST, and PUT.
	// Default: false.
	//
	HttpMethod *bool `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Represents the information of $context.requestId.
	//
	// Whether the ID for the request should be included in the request context.
	// Default: false.
	//
	RequestId *bool `field:"optional" json:"requestId" yaml:"requestId"`
	// Represents the information of $context.resourceId.
	//
	// Whether the identifier that API Gateway assigns to your resource should be included in the request context.
	// Default: false.
	//
	ResourceId *bool `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Represents the information of $context.resourcePath.
	//
	// Whether the path to the resource should be included in the request context.
	// Default: false.
	//
	ResourcePath *bool `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Represents the information of $context.identity.sourceIp.
	//
	// Whether the source IP address of the immediate TCP connection making the request
	// to API Gateway endpoint should be included in the request context.
	// Default: false.
	//
	SourceIp *bool `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Represents the information of $context.stage.
	//
	// Whether the deployment stage of the API request should be included in the request context.
	// Default: false.
	//
	Stage *bool `field:"optional" json:"stage" yaml:"stage"`
	// Represents the information of $context.identity.user.
	//
	// Whether the principal identifier of the user that will be authorized should be included in the request context.
	// Supported for resources that use IAM authorization.
	// Default: false.
	//
	User *bool `field:"optional" json:"user" yaml:"user"`
	// Represents the information of $context.identity.userAgent.
	//
	// Whether the User-Agent header of the API caller should be included in the request context.
	// Default: false.
	//
	UserAgent *bool `field:"optional" json:"userAgent" yaml:"userAgent"`
	// Represents the information of $context.identity.userArn.
	//
	// Whether the Amazon Resource Name (ARN) of the effective user identified after authentication should be included in the request context.
	// Default: false.
	//
	UserArn *bool `field:"optional" json:"userArn" yaml:"userArn"`
}

