package mixinsawsimagebuilder


// Specifies resources that lifecycle policy actions should not apply to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
type CfnLifecyclePolicyPropsMixin_ExclusionRulesProperty struct {
	// Lists configuration values that apply to AMIs that Image Builder should exclude from the lifecycle action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html#cfn-imagebuilder-lifecyclepolicy-exclusionrules-amis
	//
	Amis interface{} `field:"optional" json:"amis" yaml:"amis"`
	// Contains a list of tags that Image Builder uses to skip lifecycle actions for Image Builder image resources that have them.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-exclusionrules.html#cfn-imagebuilder-lifecyclepolicy-exclusionrules-tagmap
	//
	TagMap interface{} `field:"optional" json:"tagMap" yaml:"tagMap"`
}

