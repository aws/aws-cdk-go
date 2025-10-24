package awsappsync


// enum with all possible values for AppSync authorization type.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   api := appsync.GraphqlApi_FromGraphqlApiAttributes(this, jsii.String("ImportedAPI"), &GraphqlApiAttributes{
//   	GraphqlApiId: jsii.String("<api-id>"),
//   	GraphqlApiArn: jsii.String("<api-arn>"),
//   	GraphQLEndpointArn: jsii.String("<api-endpoint-arn>"),
//   	Visibility: appsync.Visibility_GLOBAL,
//   	Modes: []AuthorizationType{
//   		appsync.AuthorizationType_IAM,
//   	},
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Minutes(jsii.Number(1))),
//   })
//   role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("events.amazonaws.com")),
//   })
//
//   // allow EventBridge to use the `publish` mutation
//   api.GrantMutation(role, jsii.String("publish"))
//
//   rule.AddTarget(targets.NewAppSync(api, &AppSyncGraphQLApiProps{
//   	GraphQLOperation: jsii.String("mutation Publish($message: String!){ publish(message: $message) { message } }"),
//   	Variables: events.RuleTargetInput_FromObject(map[string]*string{
//   		"message": jsii.String("hello world"),
//   	}),
//   	EventRole: role,
//   }))
//
type AuthorizationType string

const (
	// API Key authorization type.
	AuthorizationType_API_KEY AuthorizationType = "API_KEY"
	// AWS IAM authorization type.
	//
	// Can be used with Cognito Identity Pool federated credentials.
	AuthorizationType_IAM AuthorizationType = "IAM"
	// Cognito User Pool authorization type.
	AuthorizationType_USER_POOL AuthorizationType = "USER_POOL"
	// OpenID Connect authorization type.
	AuthorizationType_OIDC AuthorizationType = "OIDC"
	// Lambda authorization type.
	AuthorizationType_LAMBDA AuthorizationType = "LAMBDA"
)

