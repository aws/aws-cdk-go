package awsapigateway


// Example:
//   var books resource
//   userPool := cognito.NewUserPool(this, jsii.String("UserPool"))
//
//   auth := apigateway.NewCognitoUserPoolsAuthorizer(this, jsii.String("booksAuthorizer"), &CognitoUserPoolsAuthorizerProps{
//   	CognitoUserPools: []iUserPool{
//   		userPool,
//   	},
//   })
//   books.AddMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &MethodOptions{
//   	Authorizer: auth,
//   	AuthorizationType: apigateway.AuthorizationType_COGNITO,
//   })
//
type AuthorizationType string

const (
	// Open access.
	AuthorizationType_NONE AuthorizationType = "NONE"
	// Use AWS IAM permissions.
	AuthorizationType_IAM AuthorizationType = "IAM"
	// Use a custom authorizer.
	AuthorizationType_CUSTOM AuthorizationType = "CUSTOM"
	// Use an AWS Cognito user pool.
	AuthorizationType_COGNITO AuthorizationType = "COGNITO"
)

