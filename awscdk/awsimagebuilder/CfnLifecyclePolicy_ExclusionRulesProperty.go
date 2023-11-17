package awsimagebuilder


// The exclusion rules to apply of the policy detail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   exclusionRulesProperty := &ExclusionRulesProperty{
//   	Amis: &AmiExclusionRulesProperty{
//   		IsPublic: jsii.Boolean(false),
//   		LastLaunched: &LastLaunchedProperty{
//   			Unit: jsii.String("unit"),
//   			Value: jsii.Number(123),
//   		},
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
//   		SharedAccounts: []*string{
//   			jsii.String("sharedAccounts"),
//   		},
//   		TagMap: map[string]*string{
//   			"tagMapKey": jsii.String("tagMap"),
//   		},
//   	},
//   	TagMap: map[string]*string{
//   		"tagMapKey": jsii.String("tagMap"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html
//
type CfnLifecyclePolicy_ExclusionRulesProperty struct {
	// The AMI exclusion rules for the policy detail.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html#cfn-imagebuilder-lifecyclepolicy-exclusionrules-amis
	//
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// The Image Builder tags to filter on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html#cfn-imagebuilder-lifecyclepolicy-exclusionrules-tagmap
	//
	TagMap interface{} `field:"optional" json:"tagMap" yaml:"tagMap"`
}

