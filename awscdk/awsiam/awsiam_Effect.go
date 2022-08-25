package awsiam


// The Effect element of an IAM policy.
//
// Example:
//   var books resource
//   var iamUser user
//
//
//   getBooks := books.addMethod(jsii.String("GET"), apigateway.NewHttpIntegration(jsii.String("http://amazon.com")), &methodOptions{
//   	authorizationType: apigateway.authorizationType_IAM,
//   })
//
//   iamUser.attachInlinePolicy(iam.NewPolicy(this, jsii.String("AllowBooks"), &policyProps{
//   	statements: []policyStatement{
//   		iam.NewPolicyStatement(&policyStatementProps{
//   			actions: []*string{
//   				jsii.String("execute-api:Invoke"),
//   			},
//   			effect: iam.effect_ALLOW,
//   			resources: []*string{
//   				getBooks.methodArn,
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

