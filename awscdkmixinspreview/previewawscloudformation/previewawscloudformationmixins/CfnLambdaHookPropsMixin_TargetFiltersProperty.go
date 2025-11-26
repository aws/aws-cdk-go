package previewawscloudformationmixins


// The `TargetFilters` property type specifies the target filters for the Hook.
//
// For more information, see [CloudFormation Hook target filters](https://docs.aws.amazon.com/cloudformation-cli/latest/hooks-userguide/hooks-target-filtering.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetFiltersProperty := &TargetFiltersProperty{
//   	Actions: []*string{
//   		jsii.String("actions"),
//   	},
//   	InvocationPoints: []*string{
//   		jsii.String("invocationPoints"),
//   	},
//   	TargetNames: []*string{
//   		jsii.String("targetNames"),
//   	},
//   	Targets: []interface{}{
//   		&HookTargetProperty{
//   			Action: jsii.String("action"),
//   			InvocationPoint: jsii.String("invocationPoint"),
//   			TargetName: jsii.String("targetName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-targetfilters.html
//
type CfnLambdaHookPropsMixin_TargetFiltersProperty struct {
	// List of actions that the hook is going to target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-targetfilters.html#cfn-cloudformation-lambdahook-targetfilters-actions
	//
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// List of invocation points that the hook is going to target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-targetfilters.html#cfn-cloudformation-lambdahook-targetfilters-invocationpoints
	//
	InvocationPoints *[]*string `field:"optional" json:"invocationPoints" yaml:"invocationPoints"`
	// List of type names that the hook is going to target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-targetfilters.html#cfn-cloudformation-lambdahook-targetfilters-targetnames
	//
	TargetNames *[]*string `field:"optional" json:"targetNames" yaml:"targetNames"`
	// List of hook targets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-lambdahook-targetfilters.html#cfn-cloudformation-lambdahook-targetfilters-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}

