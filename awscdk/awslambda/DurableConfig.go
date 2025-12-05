package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Configuration for durable functions.
//
// Lambda durable functions allow for long-running executions with persistent state.
//
// Example:
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_24_X(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	DurableConfig: &DurableConfig{
//   		ExecutionTimeout: awscdk.Duration_Hours(jsii.Number(1)),
//   		RetentionPeriod: awscdk.Duration_Days(jsii.Number(30)),
//   	},
//   })
//
type DurableConfig struct {
	// The amount of time that Lambda allows a durable function to run before stopping it.
	//
	// Must be between 1 and 31,622,400 seconds (366 days).
	ExecutionTimeout awscdk.Duration `field:"required" json:"executionTimeout" yaml:"executionTimeout"`
	// The number of days after a durable execution is closed that Lambda retains its history.
	//
	// Must be between 1 and 90 days.
	//
	// The underlying configuration is expressed in whole numbers of days. Providing a Duration that
	// does not represent a whole number of days will result in a runtime or deployment error.
	// Default: Duration.days(14)
	//
	RetentionPeriod awscdk.Duration `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

