package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for TokenAuthorizer.
//
// Example:
//   var authFn function
//   var books resource
//
//
//   auth := apigateway.NewTokenAuthorizer(this, jsii.String("booksAuthorizer"), &TokenAuthorizerProps{
//   	Handler: authFn,
//   })
//
//   books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
//   	Authorizer: auth,
//   })
//
type TokenAuthorizerProps struct {
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
	// The request header mapping expression for the bearer token.
	//
	// This is typically passed as part of the header, in which case
	// this should be `method.request.header.Authorizer` where `Authorizer` is the header containing the bearer token.
	// See: https://docs.aws.amazon.com/apigateway/latest/api/API_CreateAuthorizer.html#apigw-CreateAuthorizer-request-identitySource
	//
	// Default: `IdentitySource.header('Authorization')`
	//
	IdentitySource *string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// An optional regex to be matched against the authorization token.
	//
	// When matched the authorizer lambda is invoked,
	// otherwise a 401 Unauthorized is returned to the client.
	// Default: - no regex filter will be applied.
	//
	ValidationRegex *string `field:"optional" json:"validationRegex" yaml:"validationRegex"`
}

