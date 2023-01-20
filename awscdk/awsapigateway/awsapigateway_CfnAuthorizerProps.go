package awsapigateway


// Properties for defining a `CfnAuthorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAuthorizerProps := &cfnAuthorizerProps{
//   	name: jsii.String("name"),
//   	restApiId: jsii.String("restApiId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	authorizerCredentials: jsii.String("authorizerCredentials"),
//   	authorizerResultTtlInSeconds: jsii.Number(123),
//   	authorizerUri: jsii.String("authorizerUri"),
//   	authType: jsii.String("authType"),
//   	identitySource: jsii.String("identitySource"),
//   	identityValidationExpression: jsii.String("identityValidationExpression"),
//   	providerArns: []*string{
//   		jsii.String("providerArns"),
//   	},
//   }
//
type CfnAuthorizerProps struct {
	// The name of the authorizer.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the `RestApi` resource that API Gateway creates the authorizer in.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The type of authorizer. Valid values include:.
	//
	// - `TOKEN` : A custom authorizer that uses a Lambda function.
	// - `COGNITO_USER_POOLS` : An authorizer that uses Amazon Cognito user pools.
	// - `REQUEST` : An authorizer that uses a Lambda function using incoming request parameters.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The credentials that are required for the authorizer.
	//
	// To specify an IAM role that API Gateway assumes, specify the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null.
	AuthorizerCredentials *string `field:"optional" json:"authorizerCredentials" yaml:"authorizerCredentials"`
	// The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches authorizer results.
	//
	// If you specify a value greater than 0, API Gateway caches the authorizer responses. By default, API Gateway sets this property to 300. The maximum value is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// The authorizer's Uniform Resource Identifier (URI).
	//
	// If you specify `TOKEN` for the authorizer's `Type` property, specify a Lambda function URI that has the form `arn:aws:apigateway: *region* :lambda:path/ *path*` . The path usually has the form /2015-03-31/functions/ *LambdaFunctionARN* /invocations.
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// An optional customer-defined field that's used in OpenApi imports and exports without functional impact.
	AuthType *string `field:"optional" json:"authType" yaml:"authType"`
	// The source of the identity in an incoming request.
	//
	// If you specify `TOKEN` or `COGNITO_USER_POOLS` for the `Type` property, this property is required. Specify a header mapping expression using the form `method.request.header. *name*` , where *name* is the name of a custom authorization header that clients submit as part of their requests.
	//
	// If you specify `REQUEST` for the `Type` property, this property is required when authorization caching is enabled. Specify a comma-separated string of one or more mapping expressions of the specified request parameter using the form `method.request.parameter. *name*` . For supported parameter types, see [Configure Lambda Authorizer Using the API Gateway Console](https://docs.aws.amazon.com/apigateway/latest/developerguide/configure-api-gateway-lambda-authorization-with-console.html) in the *API Gateway Developer Guide* .
	IdentitySource *string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// A validation expression for the incoming identity.
	//
	// If you specify `TOKEN` for the authorizer's `Type` property, specify a regular expression. API Gateway uses the expression to attempt to match the incoming client token, and proceeds if the token matches. If the token doesn't match, API Gateway responds with a 401 (unauthorized request) error code.
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
	// A list of the Amazon Cognito user pool Amazon Resource Names (ARNs) to associate with this authorizer.
	//
	// Required if you specify `COGNITO_USER_POOLS` as the authorizer `Type` . For more information, see [Use Amazon Cognito User Pools](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html#apigateway-enable-cognito-user-pool) in the *API Gateway Developer Guide* .
	ProviderArns *[]*string `field:"optional" json:"providerArns" yaml:"providerArns"`
}

