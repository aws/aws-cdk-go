package awsdlm


// Specifies a rule for sharing snapshots across AWS accounts .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   shareRuleProperty := &shareRuleProperty{
//   	targetAccounts: []*string{
//   		jsii.String("targetAccounts"),
//   	},
//   	unshareInterval: jsii.Number(123),
//   	unshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   }
//
type CfnLifecyclePolicy_ShareRuleProperty struct {
	// The IDs of the AWS accounts with which to share the snapshots.
	TargetAccounts *[]*string `field:"optional" json:"targetAccounts" yaml:"targetAccounts"`
	// The period after which snapshots that are shared with other AWS accounts are automatically unshared.
	UnshareInterval *float64 `field:"optional" json:"unshareInterval" yaml:"unshareInterval"`
	// The unit of time for the automatic unsharing interval.
	UnshareIntervalUnit *string `field:"optional" json:"unshareIntervalUnit" yaml:"unshareIntervalUnit"`
}

