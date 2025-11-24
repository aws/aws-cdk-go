package mixinsawsbatch


// Specifies an array of up to 5 conditions to be met, and an action to take ( `RETRY` or `EXIT` ) if all conditions are met.
//
// If none of the `EvaluateOnExit` conditions in a `RetryStrategy` match, then the job is retried.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   evaluateOnExitProperty := &EvaluateOnExitProperty{
//   	Action: jsii.String("action"),
//   	OnExitCode: jsii.String("onExitCode"),
//   	OnReason: jsii.String("onReason"),
//   	OnStatusReason: jsii.String("onStatusReason"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-evaluateonexit.html
//
type CfnJobDefinitionPropsMixin_EvaluateOnExitProperty struct {
	// Specifies the action to take if all of the specified conditions ( `onStatusReason` , `onReason` , and `onExitCode` ) are met.
	//
	// The values aren't case sensitive.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-evaluateonexit.html#cfn-batch-jobdefinition-evaluateonexit-action
	//
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Contains a glob pattern to match against the decimal representation of the `ExitCode` returned for a job.
	//
	// The pattern can be up to 512 characters long. It can contain only numbers, and can end with an asterisk (*) so that only the start of the string needs to be an exact match.
	//
	// The string can contain up to 512 characters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-evaluateonexit.html#cfn-batch-jobdefinition-evaluateonexit-onexitcode
	//
	OnExitCode *string `field:"optional" json:"onExitCode" yaml:"onExitCode"`
	// Contains a glob pattern to match against the `Reason` returned for a job.
	//
	// The pattern can contain up to 512 characters. It can contain letters, numbers, periods (.), colons (:), and white space (including spaces and tabs). It can optionally end with an asterisk (*) so that only the start of the string needs to be an exact match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-evaluateonexit.html#cfn-batch-jobdefinition-evaluateonexit-onreason
	//
	OnReason *string `field:"optional" json:"onReason" yaml:"onReason"`
	// Contains a glob pattern to match against the `StatusReason` returned for a job.
	//
	// The pattern can contain up to 512 characters. It can contain letters, numbers, periods (.), colons (:), and white spaces (including spaces or tabs). It can optionally end with an asterisk (*) so that only the start of the string needs to be an exact match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-evaluateonexit.html#cfn-batch-jobdefinition-evaluateonexit-onstatusreason
	//
	OnStatusReason *string `field:"optional" json:"onStatusReason" yaml:"onStatusReason"`
}

