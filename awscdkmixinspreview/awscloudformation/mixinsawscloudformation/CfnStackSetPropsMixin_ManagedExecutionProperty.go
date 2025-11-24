package mixinsawscloudformation


// Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedExecutionProperty := &ManagedExecutionProperty{
//   	Active: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-managedexecution.html
//
type CfnStackSetPropsMixin_ManagedExecutionProperty struct {
	// When `true` , CloudFormation performs non-conflicting operations concurrently and queues conflicting operations.
	//
	// After conflicting operations finish, CloudFormation starts queued operations in request order.
	//
	// > If there are already running or queued operations, CloudFormation queues all incoming operations even if they are non-conflicting.
	// >
	// > You can't modify your StackSet's execution configuration while there are running or queued operations for that StackSet.
	//
	// When `false` (default), StackSets performs one operation at a time in request order.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-stackset-managedexecution.html#cfn-cloudformation-stackset-managedexecution-active
	//
	Active interface{} `field:"optional" json:"active" yaml:"active"`
}

