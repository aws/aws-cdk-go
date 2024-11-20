package awscloudformation


// Filters to allow hooks to target specific stack attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-stackfilters.html
//
type CfnLambdaHook_StackFiltersProperty struct {
	// Attribute to specify the filtering behavior.
	//
	// ANY will make the Hook pass if one filter matches. ALL will make the Hook pass if all filters match
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-stackfilters.html#cfn-cloudformation-lambdahook-stackfilters-filteringcriteria
	//
	// Default: - "ALL".
	//
	FilteringCriteria *string `field:"required" json:"filteringCriteria" yaml:"filteringCriteria"`
	// List of stack names as filters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-stackfilters.html#cfn-cloudformation-lambdahook-stackfilters-stacknames
	//
	StackNames interface{} `field:"optional" json:"stackNames" yaml:"stackNames"`
	// List of stack roles that are performing the stack operations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-stackfilters.html#cfn-cloudformation-lambdahook-stackfilters-stackroles
	//
	StackRoles interface{} `field:"optional" json:"stackRoles" yaml:"stackRoles"`
}

