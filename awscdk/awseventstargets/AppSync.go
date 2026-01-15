package awseventstargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents"
)

// Use an AppSync GraphQL API as a target for Amazon EventBridge rules.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("api"),
//   	Definition: appsync.Definition_FromFile(jsii.String("schema.graphql")),
//   	AuthorizationConfig: &AuthorizationConfig{
//   		DefaultAuthorization: &AuthorizationMode{
//   			AuthorizationType: appsync.AuthorizationType_IAM,
//   		},
//   	},
//   })
//
//   rule := events.NewRule(this, jsii.String("Rule"), &RuleProps{
//   	Schedule: events.Schedule_Rate(cdk.Duration_Hours(jsii.Number(1))),
//   })
//
//   rule.AddTarget(targets.NewAppSync(api, &AppSyncGraphQLApiProps{
//   	GraphQLOperation: jsii.String("mutation Publish($message: String!){ publish(message: $message) { message } }"),
//   	Variables: events.RuleTargetInput_FromObject(map[string]*string{
//   		"message": jsii.String("hello world"),
//   	}),
//   }))
//
type AppSync interface {
	awsevents.IRuleTarget
	// Returns a RuleTarget that can be used to trigger this AppSync GraphQL API as a result from an EventBridge event.
	Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig
}

// The jsii proxy struct for AppSync
type jsiiProxy_AppSync struct {
	internal.Type__awseventsIRuleTarget
}

func NewAppSync(appsyncApi awsappsync.IGraphqlApi, props *AppSyncGraphQLApiProps) AppSync {
	_init_.Initialize()

	if err := validateNewAppSyncParameters(appsyncApi, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSync{}

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.AppSync",
		[]interface{}{appsyncApi, props},
		&j,
	)

	return &j
}

func NewAppSync_Override(a AppSync, appsyncApi awsappsync.IGraphqlApi, props *AppSyncGraphQLApiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_events_targets.AppSync",
		[]interface{}{appsyncApi, props},
		a,
	)
}

func (a *jsiiProxy_AppSync) Bind(rule interfacesawsevents.IRuleRef, id *string) *awsevents.RuleTargetConfig {
	if err := a.validateBindParameters(rule); err != nil {
		panic(err)
	}
	var returns *awsevents.RuleTargetConfig

	_jsii_.Invoke(
		a,
		"bind",
		[]interface{}{rule, id},
		&returns,
	)

	return returns
}

