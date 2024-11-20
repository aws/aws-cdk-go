package awscloudformation


// List of stack names as filters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stacknames.html
//
type CfnGuardHook_StackNamesProperty struct {
	// List of stack names that the hook is going to be excluded from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stacknames.html#cfn-cloudformation-guardhook-stacknames-exclude
	//
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// List of stack names that the hook is going to target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudformation-guardhook-stacknames.html#cfn-cloudformation-guardhook-stacknames-include
	//
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

