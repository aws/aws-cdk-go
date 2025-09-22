package awscloudformation


// Hook targets are the destination where hooks will be invoked against.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hookTargetProperty := &HookTargetProperty{
//   	Action: jsii.String("action"),
//   	InvocationPoint: jsii.String("invocationPoint"),
//   	TargetName: jsii.String("targetName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-hooktarget.html
//
type CfnGuardHook_HookTargetProperty struct {
	// Target actions are the type of operation hooks will be executed at.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-hooktarget.html#cfn-cloudformation-guardhook-hooktarget-action
	//
	Action *string `field:"required" json:"action" yaml:"action"`
	// Invocation points are the point in provisioning workflow where hooks will be executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-hooktarget.html#cfn-cloudformation-guardhook-hooktarget-invocationpoint
	//
	InvocationPoint *string `field:"required" json:"invocationPoint" yaml:"invocationPoint"`
	// Type name of hook target.
	//
	// Hook targets are the destination where hooks will be invoked against.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-hooktarget.html#cfn-cloudformation-guardhook-hooktarget-targetname
	//
	TargetName *string `field:"required" json:"targetName" yaml:"targetName"`
}

