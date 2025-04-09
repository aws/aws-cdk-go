package awscdk


// The `StackFilters` property type specifies stack level filters for a Hook.
//
// The `StackNames` or `StackRoles` properties are optional. However, you must specify at least one of these properties.
//
// For more information, see [AWS CloudFormation Hooks stack level filters](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-stack-level-filtering.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   stackFiltersProperty := &StackFiltersProperty{
//   	FilteringCriteria: jsii.String("filteringCriteria"),
//
//   	// the properties below are optional
//   	StackNames: &StackNamesProperty{
//   		Exclude: []*string{
//   			jsii.String("exclude"),
//   		},
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   	StackRoles: &StackRolesProperty{
//   		Exclude: []*string{
//   			jsii.String("exclude"),
//   		},
//   		Include: []*string{
//   			jsii.String("include"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackfilters.html
//
type CfnGuardHook_StackFiltersProperty struct {
	// The filtering criteria.
	//
	// - All stack names and stack roles ( `All` ): The Hook will only be invoked when all specified filters match.
	// - Any stack names and stack roles ( `Any` ): The Hook will be invoked if at least one of the specified filters match.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackfilters.html#cfn-cloudformation-guardhook-stackfilters-filteringcriteria
	//
	// Default: - "ALL".
	//
	FilteringCriteria *string `field:"required" json:"filteringCriteria" yaml:"filteringCriteria"`
	// Includes or excludes specific stacks from Hook invocations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackfilters.html#cfn-cloudformation-guardhook-stackfilters-stacknames
	//
	StackNames interface{} `field:"optional" json:"stackNames" yaml:"stackNames"`
	// Includes or excludes specific stacks from Hook invocations based on their associated IAM roles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackfilters.html#cfn-cloudformation-guardhook-stackfilters-stackroles
	//
	StackRoles interface{} `field:"optional" json:"stackRoles" yaml:"stackRoles"`
}

