package previewawsimagebuildermixins


// Defines criteria for AMIs that are excluded from lifecycle actions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnLifecyclePolicyPropsMixin_AmiExclusionRulesProperty struct {
	// Configures whether public AMIs are excluded from the lifecycle action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-ispublic
	//
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// Specifies configuration details for Image Builder to exclude the most recent resources from lifecycle actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-lastlaunched
	//
	LastLaunched interface{} `field:"optional" json:"lastLaunched" yaml:"lastLaunched"`
	// Configures AWS Region s that are excluded from the lifecycle action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Specifies AWS account s whose resources are excluded from the lifecycle action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-sharedaccounts
	//
	SharedAccounts *[]*string `field:"optional" json:"sharedAccounts" yaml:"sharedAccounts"`
	// Lists tags that should be excluded from lifecycle actions for the AMIs that have them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-amiexclusionrules.html#cfn-imagebuilder-lifecyclepolicy-amiexclusionrules-tagmap
	//
	TagMap interface{} `field:"optional" json:"tagMap" yaml:"tagMap"`
}

