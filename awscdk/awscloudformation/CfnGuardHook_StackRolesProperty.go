package awscloudformation


// Specifies the stack roles for the `StackFilters` property type to include or exclude specific stacks from Hook invocations based on their associated IAM roles.
//
// For more information, see [AWS CloudFormation Hooks stack level filters](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-stack-level-filtering.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackRolesProperty := &StackRolesProperty{
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	Include: []*string{
//   		jsii.String("include"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackroles.html
//
type CfnGuardHook_StackRolesProperty struct {
	// The IAM role ARNs for stacks you want to exclude.
	//
	// The Hook will be invoked on all stacks except those initiated by the specified roles.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackroles.html#cfn-cloudformation-guardhook-stackroles-exclude
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// The IAM role ARNs to target stacks associated with these roles.
	//
	// Only stack operations initiated by these roles will invoke the Hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackroles.html#cfn-cloudformation-guardhook-stackroles-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

