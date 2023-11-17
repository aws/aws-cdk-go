package awsimagebuilder


// The AMI exclusion rules for the policy detail.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   amiExclusionRulesProperty := &AmiExclusionRulesProperty{
//   	IsPublic: jsii.Boolean(false),
//   	LastLaunched: &LastLaunchedProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	SharedAccounts: []*string{
//   		jsii.String("sharedAccounts"),
//   	},
//   	TagMap: map[string]*string{
//   		"tagMapKey": jsii.String("tagMap"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html
//
type CfnLifecyclePolicy_AmiExclusionRulesProperty struct {
	// Use to apply lifecycle policy actions on whether the AMI is public.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-ispublic
	//
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// The last launched time of a resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-lastlaunched
	//
	LastLaunched interface{} `field:"optional" json:"lastLaunched" yaml:"lastLaunched"`
	// Use to apply lifecycle policy actions on AMIs distributed to a set of regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Use to apply lifecycle policy actions on AMIs shared with a set of regions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-sharedaccounts
	//
	SharedAccounts *[]*string `field:"optional" json:"sharedAccounts" yaml:"sharedAccounts"`
	// The AMIs to select by tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-tagmap
	//
	TagMap interface{} `field:"optional" json:"tagMap" yaml:"tagMap"`
}

