package awsiam


// The Effect element of an IAM policy.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // This function handles your connect route
//   var connectHandler Function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"))
//
//   webSocketApi.AddRoute(jsii.String("$connect"), &WebSocketRouteOptions{
//   	Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("Integration"), connectHandler),
//   	Authorizer: awscdk.NewWebSocketIamAuthorizer(),
//   })
//
//   // Create an IAM user (identity)
//   user := iam.NewUser(this, jsii.String("User"))
//
//   webSocketArn := awscdk.stack_Of(this).FormatArn(&ArnComponents{
//   	Service: jsii.String("execute-api"),
//   	Resource: webSocketApi.ApiId,
//   })
//
//   // Grant access to the IAM user
//   user.AttachInlinePolicy(iam.NewPolicy(this, jsii.String("AllowInvoke"), &PolicyProps{
//   	Statements: []PolicyStatement{
//   		iam.NewPolicyStatement(&PolicyStatementProps{
//   			Actions: []*string{
//   				jsii.String("execute-api:Invoke"),
//   			},
//   			Effect: iam.Effect_ALLOW,
//   			Resources: []*string{
//   				webSocketArn,
//   			},
//   		}),
//   	},
//   }))
//
// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_effect.html
//
type Effect string

const (
	// Allows access to a resource in an IAM policy statement.
	//
	// By default, access to resources are denied.
	Effect_ALLOW Effect = "ALLOW"
	// Explicitly deny access to a resource.
	//
	// By default, all requests are denied implicitly.
	// See: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html
	//
	Effect_DENY Effect = "DENY"
)

