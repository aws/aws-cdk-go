package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for RequestAuthorizer.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewRequestAuthorizer(this, jsii.String("booksAuthorizer"), &RequestAuthorizerProps{
//   	Handler: authFn,
//   	IdentitySources: []*string{
//   		apigateway.IdentitySource_Header(jsii.String("Authorization")),
//   	},
//   })
//
//   books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
//   	Authorizer: auth,
//   })
//
type RequestAuthorizerProps struct {
	// The handler for the authorizer lambda function.
	//
	// The handler must follow a very specific protocol on the input it receives
	// and the output it needs to produce.  API Gateway has documented the
	// handler's [input specification](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-input.html)
	// and [output specification](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-output.html).
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// An optional IAM role for APIGateway to assume before calling the Lambda-based authorizer.
	//
	// The IAM role must be
	// assumable by 'apigateway.amazonaws.com'.
	// Default: - A resource policy is added to the Lambda function allowing apigateway.amazonaws.com to invoke the function.
	//
	AssumeRole awsiam.IRole `field:"optional" json:"assumeRole" yaml:"assumeRole"`
	// An optional human friendly name for the authorizer.
	//
	// Note that, this is not the primary identifier of the authorizer.
	// Default: - the unique construct ID.
	//
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to 0.
	// Default: - Duration.minutes(5)
	//
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
	// An array of request header mapping expressions for identities.
	//
	// Supported parameter types are
	// Header, Query String, Stage Variable, and Context. For instance, extracting an authorization
	// token from a header would use the identity source `IdentitySource.header('Authorization')`.
	//
	// Note: API Gateway uses the specified identity sources as the request authorizer caching key. When caching is
	// enabled, API Gateway calls the authorizer's Lambda function only after successfully verifying that all the
	// specified identity sources are present at runtime. If a specified identify source is missing, null, or empty,
	// API Gateway returns a 401 Unauthorized response without calling the authorizer Lambda function.
	// See: https://docs.aws.amazon.com/apigateway/latest/api/API_CreateAuthorizer.html#apigw-CreateAuthorizer-request-identitySource
	//
	IdentitySources *[]*string `field:"required" json:"identitySources" yaml:"identitySources"`
}

