package triggers

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// Props for `Trigger`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   func := lambda.NewFunction(this, jsii.String("MyFunction"), &functionProps{
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   	Code: lambda.Code_FromInline(jsii.String("foo")),
//   })
//
//   triggers.NewTrigger(this, jsii.String("MyTrigger"), &TriggerProps{
//   	Handler: func,
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   	InvocationType: triggers.InvocationType_EVENT,
//   })
//
type TriggerProps struct {
	// Adds trigger dependencies. Execute this trigger only after these construct scopes have been provisioned.
	//
	// You can also use `trigger.executeAfter()` to add additional dependencies.
	// Default: [].
	//
	ExecuteAfter *[]constructs.Construct `field:"optional" json:"executeAfter" yaml:"executeAfter"`
	// Adds this trigger as a dependency on other constructs.
	//
	// This means that this
	// trigger will get executed *before* the given construct(s).
	//
	// You can also use `trigger.executeBefore()` to add additional dependants.
	// Default: [].
	//
	ExecuteBefore *[]constructs.Construct `field:"optional" json:"executeBefore" yaml:"executeBefore"`
	// Re-executes the trigger every time the handler changes.
	//
	// This implies that the trigger is associated with the `currentVersion` of
	// the handler, which gets recreated every time the handler or its
	// configuration is updated.
	// Default: true.
	//
	ExecuteOnHandlerChange *bool `field:"optional" json:"executeOnHandlerChange" yaml:"executeOnHandlerChange"`
	// The AWS Lambda function of the handler to execute.
	Handler awslambda.Function `field:"required" json:"handler" yaml:"handler"`
	// The invocation type to invoke the Lambda function with.
	// Default: RequestResponse.
	//
	InvocationType InvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// The timeout of the invocation call of the Lambda function to be triggered.
	// Default: Duration.minutes(2)
	//
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}

