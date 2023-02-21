package awsapigateway


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
type AuthorizationType string

const (
	// Open access.
	// Experimental.
	AuthorizationType_NONE AuthorizationType = "NONE"
	// Use AWS IAM permissions.
	// Experimental.
	AuthorizationType_IAM AuthorizationType = "IAM"
	// Use a custom authorizer.
	// Experimental.
	AuthorizationType_CUSTOM AuthorizationType = "CUSTOM"
	// Use an AWS Cognito user pool.
	// Experimental.
	AuthorizationType_COGNITO AuthorizationType = "COGNITO"
)

