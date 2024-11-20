package awscdk


// Properties for defining a `CfnLambdaHook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
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
	// The typename alias for the hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// IAM Role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Attribute to specify CloudFormation behavior on hook failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-failuremode
	//
	FailureMode *string `field:"required" json:"failureMode" yaml:"failureMode"`
	// Attribute to specify which stacks this hook applies to or should get invoked for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-hookstatus
	//
	// Default: - "ENABLED".
	//
	HookStatus *string `field:"required" json:"hookStatus" yaml:"hookStatus"`
	// Amazon Resource Name (ARN), Partial ARN, name, version, or alias of the Lambda function to invoke with this hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-lambdafunction
	//
	LambdaFunction *string `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Which operations should this Hook run against?
	//
	// Resource changes, stacks or change sets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-targetoperations
	//
	TargetOperations *[]*string `field:"required" json:"targetOperations" yaml:"targetOperations"`
	// Filters to allow hooks to target specific stack attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-stackfilters
	//
	StackFilters interface{} `field:"optional" json:"stackFilters" yaml:"stackFilters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-lambdahook.html#cfn-cloudformation-lambdahook-targetfilters
	//
	TargetFilters interface{} `field:"optional" json:"targetFilters" yaml:"targetFilters"`
}

