package awsapigatewayv2


// Properties for defining a `CfnAuthorizer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAuthorizerProps := &CfnAuthorizerProps{
//   	ApiId: jsii.String("apiId"),
//   	AuthorizerType: jsii.String("authorizerType"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AuthorizerCredentialsArn: jsii.String("authorizerCredentialsArn"),
//   	AuthorizerPayloadFormatVersion: jsii.String("authorizerPayloadFormatVersion"),
//   	AuthorizerResultTtlInSeconds: jsii.Number(123),
//   	AuthorizerUri: jsii.String("authorizerUri"),
//   	EnableSimpleResponses: jsii.Boolean(false),
//   	IdentitySource: []*string{
//   		jsii.String("identitySource"),
//   	},
//   	IdentityValidationExpression: jsii.String("identityValidationExpression"),
//   	JwtConfiguration: &JWTConfigurationProperty{
//   		Audience: []*string{
//   			jsii.String("audience"),
//   		},
//   		Issuer: jsii.String("issuer"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html
//
type CfnAuthorizerProps struct {
	// The API identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The authorizer type.
	//
	// Specify `REQUEST` for a Lambda function using incoming request parameters. Specify `JWT` to use JSON Web Tokens (supported only for HTTP APIs).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizertype
	//
	AuthorizerType *string `field:"required" json:"authorizerType" yaml:"authorizerType"`
	// The name of the authorizer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer.
	//
	// To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for `REQUEST` authorizers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizercredentialsarn
	//
	AuthorizerCredentialsArn *string `field:"optional" json:"authorizerCredentialsArn" yaml:"authorizerCredentialsArn"`
	// Specifies the format of the payload sent to an HTTP API Lambda authorizer.
	//
	// Required for HTTP API Lambda authorizers. Supported values are `1.0` and `2.0` . To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizerpayloadformatversion
	//
	AuthorizerPayloadFormatVersion *string `field:"optional" json:"authorizerPayloadFormatVersion" yaml:"authorizerPayloadFormatVersion"`
	// The time to live (TTL) for cached authorizer results, in seconds.
	//
	// If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizerresultttlinseconds
	//
	AuthorizerResultTtlInSeconds *float64 `field:"optional" json:"authorizerResultTtlInSeconds" yaml:"authorizerResultTtlInSeconds"`
	// The authorizer's Uniform Resource Identifier (URI).
	//
	// For `REQUEST` authorizers, this must be a well-formed Lambda function URI, for example, `arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2: *{account_id}* :function: *{lambda_function_name}* /invocations` . In general, the URI has this form: `arn:aws:apigateway: *{region}* :lambda:path/ *{service_api}*` , where *{region}* is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial `/` . For Lambda functions, this is usually of the form `/2015-03-31/functions/[FunctionARN]/invocations` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-authorizeruri
	//
	AuthorizerUri *string `field:"optional" json:"authorizerUri" yaml:"authorizerUri"`
	// Specifies whether a Lambda authorizer returns a response in a simple format.
	//
	// By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-enablesimpleresponses
	//
	EnableSimpleResponses interface{} `field:"optional" json:"enableSimpleResponses" yaml:"enableSimpleResponses"`
	// The identity source for which authorization is requested.
	//
	// For a `REQUEST` authorizer, this is optional. The value is a set of one or more mapping expressions of the specified request parameters. The identity source can be headers, query string parameters, stage variables, and context parameters. For example, if an Auth header and a Name query string parameter are defined as identity sources, this value is route.request.header.Auth, route.request.querystring.Name for WebSocket APIs. For HTTP APIs, use selection expressions prefixed with `$` , for example, `$request.header.Auth` , `$request.querystring.Name` . These parameters are used to perform runtime validation for Lambda-based authorizers by verifying all of the identity-related request parameters are present in the request, not null, and non-empty. Only when this is true does the authorizer invoke the authorizer Lambda function. Otherwise, it returns a 401 Unauthorized response without calling the Lambda function. For HTTP APIs, identity sources are also used as the cache key when caching is enabled. To learn more, see [Working with AWS Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html) .
	//
	// For `JWT` , a single entry that specifies where to extract the JSON Web Token (JWT) from inbound requests. Currently only header-based and query parameter-based selections are supported, for example `$request.header.Authorization` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-identitysource
	//
	IdentitySource *[]*string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// This parameter is not used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-identityvalidationexpression
	//
	IdentityValidationExpression *string `field:"optional" json:"identityValidationExpression" yaml:"identityValidationExpression"`
	// The `JWTConfiguration` property specifies the configuration of a JWT authorizer.
	//
	// Required for the `JWT` authorizer type. Supported only for HTTP APIs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-authorizer.html#cfn-apigatewayv2-authorizer-jwtconfiguration
	//
	JwtConfiguration interface{} `field:"optional" json:"jwtConfiguration" yaml:"jwtConfiguration"`
}

