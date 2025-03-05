package awseventstargets

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Customize the AppSync GraphQL API target.
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
type AppSyncGraphQLApiProps struct {
	// The SQS queue to be used as deadLetterQueue. Check out the [considerations for using a dead-letter queue](https://docs.aws.amazon.com/eventbridge/latest/userguide/rule-dlq.html#dlq-considerations).
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The maximum age of a request that Lambda sends to a function for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the function returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The GraphQL operation;
	//
	// that is, the query, mutation, or subscription
	// to be parsed and executed by the GraphQL service.
	GraphQLOperation *string `field:"required" json:"graphQLOperation" yaml:"graphQLOperation"`
	// The role to assume before invoking the target (i.e., the pipeline) when the given rule is triggered.
	// Default: - a new role with permissions to access mutations will be created.
	//
	EventRole awsiam.IRole `field:"optional" json:"eventRole" yaml:"eventRole"`
	// The variables that are include in the GraphQL operation.
	// Default: - The entire event is used.
	//
	Variables awsevents.RuleTargetInput `field:"optional" json:"variables" yaml:"variables"`
}

