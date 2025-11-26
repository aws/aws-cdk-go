package previewawscloudformationmixins


// Specifies the stack names for the `StackFilters` property type to include or exclude specific stacks from Hook invocations.
//
// For more information, see [CloudFormation Hooks stack level filters](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-stack-level-filtering.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stackNamesProperty := &StackNamesProperty{
//   	Exclude: []*string{
//   		jsii.String("exclude"),
//   	},
//   	Include: []*string{
//   		jsii.String("include"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-stacknames.html
//
type CfnLambdaHookPropsMixin_StackNamesProperty struct {
	// The stack names to exclude.
	//
	// All stacks except those listed here will invoke the Hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-stacknames.html#cfn-cloudformation-lambdahook-stacknames-exclude
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// The stack names to include.
	//
	// Only the stacks specified in this list will invoke the Hook.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-stacknames.html#cfn-cloudformation-lambdahook-stacknames-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

