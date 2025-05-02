package awscdk


// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   managedExecutionProperty := &ManagedExecutionProperty{
//   	Active: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-managedexecution.html
//
type CfnStackSet_ManagedExecutionProperty struct {
	// When `true` , StackSets performs non-conflicting operations concurrently and queues conflicting operations.
	//
	// After conflicting operations finish, StackSets starts queued operations in request order.
	//
	// > If there are already running or queued operations, StackSets queues all incoming operations even if they are non-conflicting.
	// >
	// > You can't modify your stack set's execution configuration while there are running or queued operations for that stack set.
	//
	// When `false` (default), StackSets performs one operation at a time in request order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-managedexecution.html#cfn-cloudformation-stackset-managedexecution-active
	//
	Active interface{} `field:"optional" json:"active" yaml:"active"`
}

