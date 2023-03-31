package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscognito"
)

// Properties for CognitoUserPoolsAuthorizer.
//
// Example:
//   var books resource
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   auth := apigateway.NewCognitoUserPoolsAuthorizer(this, jsii.String("booksAuthorizer"), &cognitoUserPoolsAuthorizerProps{
//   	cognitoUserPools: []iUserPool{
//   		userPool,
//   	},
//   })
//   books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizer: auth,
//   	authorizationType: apigateway.authorizationType_COGNITO,
//   })
//
// Experimental.
type CognitoUserPoolsAuthorizerProps struct {
	// The user pools to associate with this authorizer.
	// Experimental.
	CognitoUserPools *[]awscognito.IUserPool `field:"required" json:"cognitoUserPools" yaml:"cognitoUserPools"`
	// An optional human friendly name for the authorizer.
	//
	// Note that, this is not the primary identifier of the authorizer.
	// Experimental.
	AuthorizerName *string `field:"optional" json:"authorizerName" yaml:"authorizerName"`
	// The request header mapping expression for the bearer token.
	//
	// This is typically passed as part of the header, in which case
	// this should be `method.request.header.Authorizer` where Authorizer is the header containing the bearer token.
	// See: https://docs.aws.amazon.com/apigateway/api-reference/link-relation/authorizer-create/#identitySource
	//
	// Experimental.
	IdentitySource *string `field:"optional" json:"identitySource" yaml:"identitySource"`
	// How long APIGateway should cache the results.
	//
	// Max 1 hour.
	// Disable caching by setting this to 0.
	// Experimental.
	ResultsCacheTtl awscdk.Duration `field:"optional" json:"resultsCacheTtl" yaml:"resultsCacheTtl"`
}

