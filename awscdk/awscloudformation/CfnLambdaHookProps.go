package awscloudformation


// Properties for defining a `CfnLambdaHook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLambdaHookProps := &CfnLambdaHookProps{
//   	Alias: jsii.String("alias"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	FailureMode: jsii.String("failureMode"),
//   	HookStatus: jsii.String("hookStatus"),
//   	LambdaFunction: jsii.String("lambdaFunction"),
//   	TargetOperations: []*string{
//   		jsii.String("targetOperations"),
//   	},
//
//   	// the properties below are optional
//   	StackFilters: &StackFiltersProperty{
//   		FilteringCriteria: jsii.String("filteringCriteria"),
//
//   		// the properties below are optional
//   		StackNames: &StackNamesProperty{
//   			Exclude: []*string{
//   				jsii.String("exclude"),
//   			},
//   			Include: []*string{
//   				jsii.String("include"),
//   			},
//   		},
//   		StackRoles: &StackRolesProperty{
//   			Exclude: []*string{
//   				jsii.String("exclude"),
//   			},
//   			Include: []*string{
//   				jsii.String("include"),
//   			},
//   		},
//   	},
//   	TargetFilters: &TargetFiltersProperty{
//   		Targets: []interface{}{
//   			&HookTargetProperty{
//   				Action: jsii.String("action"),
//   				InvocationPoint: jsii.String("invocationPoint"),
//   				TargetName: jsii.String("targetName"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Actions: []*string{
//   			jsii.String("actions"),
//   		},
//   		InvocationPoints: []*string{
//   			jsii.String("invocationPoints"),
//   		},
//   		TargetNames: []*string{
//   			jsii.String("targetNames"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html
//
type CfnLambdaHookProps struct {
	// The type name alias for the Hook. This alias must be unique per account and Region.
	//
	// The alias must be in the form `Name1::Name2::Name3` and must not begin with `AWS` . For example, `Private::Lambda::MyTestHook` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The IAM role that the Hook assumes to invoke your Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Specifies how the Hook responds when the Lambda function invoked by the Hook returns a `FAILED` response.
	//
	// - `FAIL` : Prevents the action from proceeding. This is helpful for enforcing strict compliance or security policies.
	// - `WARN` : Issues warnings to users but allows actions to continue. This is useful for non-critical validations or informational checks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-failuremode
	//
	FailureMode *string `field:"required" json:"failureMode" yaml:"failureMode"`
	// Specifies if the Hook is `ENABLED` or `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-hookstatus
	//
	// Default: - "ENABLED".
	//
	HookStatus *string `field:"required" json:"hookStatus" yaml:"hookStatus"`
	// Specifies the Lambda function for the Hook. You can use:.
	//
	// - The full Amazon Resource Name (ARN) without a suffix.
	// - A qualified ARN with a version or alias suffix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-lambdafunction
	//
	LambdaFunction *string `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Specifies the list of operations the Hook is run against.
	//
	// For more information, see [Hook targets](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-concepts.html#hook-terms-hook-target) in the *AWS CloudFormation Hooks User Guide* .
	//
	// Valid values: `STACK` | `RESOURCE` | `CHANGE_SET` | `CLOUD_CONTROL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-targetoperations
	//
	TargetOperations *[]*string `field:"required" json:"targetOperations" yaml:"targetOperations"`
	// Specifies the stack level filters for the Hook.
	//
	// Example stack level filter in JSON:
	//
	// `"StackFilters": {"FilteringCriteria": "ALL", "StackNames": {"Exclude": [ "stack-1", "stack-2"]}}`
	//
	// Example stack level filter in YAML:
	//
	// `StackFilters: FilteringCriteria: ALL StackNames: Exclude: - stack-1 - stack-2`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-stackfilters
	//
	StackFilters interface{} `field:"optional" json:"stackFilters" yaml:"stackFilters"`
	// Specifies the target filters for the Hook.
	//
	// Example target filter in JSON:
	//
	// `"TargetFilters": {"Actions": [ "CREATE", "UPDATE", "DELETE" ]}`
	//
	// Example target filter in YAML:
	//
	// `TargetFilters: Actions: - CREATE - UPDATE - DELETE`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-targetfilters
	//
	TargetFilters interface{} `field:"optional" json:"targetFilters" yaml:"targetFilters"`
}

