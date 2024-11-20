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
	// The typename alias for the hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-alias
	//
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// IAM Role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// Attribute to specify CloudFormation behavior on hook failure.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-failuremode
	//
	// Default: - "WARN".
	//
	FailureMode *string `field:"required" json:"failureMode" yaml:"failureMode"`
	// Attribute to specify which stacks this hook applies to or should get invoked for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-hookstatus
	//
	// Default: - "DISABLED".
	//
	HookStatus *string `field:"required" json:"hookStatus" yaml:"hookStatus"`
	// S3 Source Location for the Guard files.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-rulelocation
	//
	RuleLocation interface{} `field:"required" json:"ruleLocation" yaml:"ruleLocation"`
	// Which operations should this Hook run against?
	//
	// Resource changes, stacks or change sets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-targetoperations
	//
	TargetOperations *[]*string `field:"required" json:"targetOperations" yaml:"targetOperations"`
	// S3 Bucket where the guard validate report will be uploaded to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-logbucket
	//
	LogBucket *string `field:"optional" json:"logBucket" yaml:"logBucket"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// Filters to allow hooks to target specific stack attributes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-stackfilters
	//
	StackFilters interface{} `field:"optional" json:"stackFilters" yaml:"stackFilters"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-guardhook.html#cfn-cloudformation-guardhook-targetfilters
	//
	TargetFilters interface{} `field:"optional" json:"targetFilters" yaml:"targetFilters"`
}

