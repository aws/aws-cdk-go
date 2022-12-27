package awsbatch


// Specifies a set of conditions to be met, and an action to take ( `RETRY` or `EXIT` ) if all conditions are met.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   evaluateOnExitProperty := &evaluateOnExitProperty{
//   	action: jsii.String("action"),
//
//   	// the properties below are optional
//   	onExitCode: jsii.String("onExitCode"),
//   	onReason: jsii.String("onReason"),
//   	onStatusReason: jsii.String("onStatusReason"),
//   }
//
type CfnJobDefinition_EvaluateOnExitProperty struct {
	// Specifies the action to take if all of the specified conditions ( `onStatusReason` , `onReason` , and `onExitCode` ) are met.
	//
	// The values aren't case sensitive.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Contains a glob pattern to match against the decimal representation of the `ExitCode` returned for a job.
	//
	// The pattern can be up to 512 characters in length. It can contain only numbers, and can optionally end with an asterisk (*) so that only the start of the string needs to be an exact match.
	//
	// The string can be between 1 and 512 characters in length.
	OnExitCode *string `field:"optional" json:"onExitCode" yaml:"onExitCode"`
	// Contains a glob pattern to match against the `Reason` returned for a job.
	//
	// The pattern can be up to 512 characters in length. It can contain letters, numbers, periods (.), colons (:), and white space (including spaces and tabs). It can optionally end with an asterisk (*) so that only the start of the string needs to be an exact match.
	//
	// The string can be between 1 and 512 characters in length.
	OnReason *string `field:"optional" json:"onReason" yaml:"onReason"`
	// Contains a glob pattern to match against the `StatusReason` returned for a job.
	//
	// The pattern can be up to 512 characters in length. It can contain letters, numbers, periods (.), colons (:), and white space (including spaces or tabs). It can optionally end with an asterisk (*) so that only the start of the string needs to be an exact match.
	//
	// The string can be between 1 and 512 characters in length.
	OnStatusReason *string `field:"optional" json:"onStatusReason" yaml:"onStatusReason"`
}

