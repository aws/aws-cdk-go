package awscloudformation


// List of stack roles that are performing the stack operations.
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
	// List of stack roles that the hook is going to be excluded from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackroles.html#cfn-cloudformation-guardhook-stackroles-exclude
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of stack roles that the hook is going to target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stackroles.html#cfn-cloudformation-guardhook-stackroles-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

