package awscdk


// Properties for defining a `CfnGuardHook`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGuardHookProps := &CfnGuardHookProps{
//   	Alias: jsii.String("alias"),
//   	ExecutionRole: jsii.String("executionRole"),
//   	FailureMode: jsii.String("failureMode"),
//   	HookStatus: jsii.String("hookStatus"),
//   	RuleLocation: &S3LocationProperty{
//   		Uri: jsii.String("uri"),
//
//   		// the properties below are optional
//   		VersionId: jsii.String("versionId"),
//   	},
//   	TargetOperations: []*string{
//   		jsii.String("targetOperations"),
//   	},
//
//   	// the properties below are optional
//   	LogBucket: jsii.String("logBucket"),
//   	Options: &OptionsProperty{
//   		InputParams: &S3LocationProperty{
//   			Uri: jsii.String("uri"),
//
//   			// the properties below are optional
//   			VersionId: jsii.String("versionId"),
//   		},
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html
//
type CfnGuardHookProps struct {
	// The type name alias for the Hook. This alias must be unique per account and Region.
	//
	// The alias must be in the form `Name1::Name2::Name3` and must not begin with `AWS` . For example, `Private::Guard::MyTestHook` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The IAM role that the Hook assumes to retrieve your Guard rules from S3 and optionally write a detailed Guard output report back.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Specifies how the Hook responds when rules fail their evaluation.
	//
	// - `FAIL` : Prevents the action from proceeding. This is helpful for enforcing strict compliance or security policies.
	// - `WARN` : Issues warnings to users but allows actions to continue. This is useful for non-critical validations or informational checks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-failuremode
	//
	// Default: - "WARN".
	//
	FailureMode *string `field:"required" json:"failureMode" yaml:"failureMode"`
	// Specifies if the Hook is `ENABLED` or `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-hookstatus
	//
	// Default: - "DISABLED".
	//
	HookStatus *string `field:"required" json:"hookStatus" yaml:"hookStatus"`
	// Specifies the S3 location of your Guard rules.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-rulelocation
	//
	RuleLocation interface{} `field:"required" json:"ruleLocation" yaml:"ruleLocation"`
	// Specifies the list of operations the Hook is run against.
	//
	// For more information, see [Hook targets](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-concepts.html#hook-terms-hook-target) in the *AWS CloudFormation Hooks User Guide* .
	//
	// Valid values: `STACK` | `RESOURCE` | `CHANGE_SET` | `CLOUD_CONTROL`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-targetoperations
	//
	TargetOperations *[]*string `field:"required" json:"targetOperations" yaml:"targetOperations"`
	// Specifies the name of an S3 bucket to store the Guard output report.
	//
	// This report contains the results of your Guard rule validations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-logbucket
	//
	LogBucket *string `field:"optional" json:"logBucket" yaml:"logBucket"`
	// Specifies the S3 location of your input parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Specifies the stack level filters for the Hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-stackfilters
	//
	StackFilters interface{} `field:"optional" json:"stackFilters" yaml:"stackFilters"`
	// Specifies the target filters for the Hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-targetfilters
	//
	TargetFilters interface{} `field:"optional" json:"targetFilters" yaml:"targetFilters"`
}

