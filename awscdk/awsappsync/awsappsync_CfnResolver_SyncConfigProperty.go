package awsappsync


// Describes a Sync configuration for a resolver.
//
// Specifies which Conflict Detection strategy and Resolution strategy to use when the resolver is invoked.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syncConfigProperty := &syncConfigProperty{
//   	conflictDetection: jsii.String("conflictDetection"),
//
//   	// the properties below are optional
//   	conflictHandler: jsii.String("conflictHandler"),
//   	lambdaConflictHandlerConfig: &lambdaConflictHandlerConfigProperty{
//   		lambdaConflictHandlerArn: jsii.String("lambdaConflictHandlerArn"),
//   	},
//   }
//
type CfnResolver_SyncConfigProperty struct {
	// The Conflict Detection strategy to use.
	//
	// - *VERSION* : Detect conflicts based on object versions for this resolver.
	// - *NONE* : Do not detect conflicts when invoking this resolver.
	ConflictDetection *string `field:"required" json:"conflictDetection" yaml:"conflictDetection"`
	// The Conflict Resolution strategy to perform in the event of a conflict.
	//
	// - *OPTIMISTIC_CONCURRENCY* : Resolve conflicts by rejecting mutations when versions don't match the latest version at the server.
	// - *AUTOMERGE* : Resolve conflicts with the Automerge conflict resolution strategy.
	// - *LAMBDA* : Resolve conflicts with an AWS Lambda function supplied in the `LambdaConflictHandlerConfig` .
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// The `LambdaConflictHandlerConfig` when configuring `LAMBDA` as the Conflict Handler.
	LambdaConflictHandlerConfig interface{} `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

