package previewawsimagebuildermixins


// The configuration details for a lifecycle policy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   policyDetailProperty := &PolicyDetailProperty{
//   	Action: &ActionProperty{
//   		IncludeResources: &IncludeResourcesProperty{
//   			Amis: jsii.Boolean(false),
//   			Containers: jsii.Boolean(false),
//   			Snapshots: jsii.Boolean(false),
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	ExclusionRules: &ExclusionRulesProperty{
//   		Amis: &AmiExclusionRulesProperty{
//   			IsPublic: jsii.Boolean(false),
//   			LastLaunched: &LastLaunchedProperty{
//   				Unit: jsii.String("unit"),
//   				Value: jsii.Number(123),
//   			},
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   			SharedAccounts: []*string{
//   				jsii.String("sharedAccounts"),
//   			},
//   			TagMap: map[string]*string{
//   				"tagMapKey": jsii.String("tagMap"),
//   			},
//   		},
//   		TagMap: map[string]*string{
//   			"tagMapKey": jsii.String("tagMap"),
//   		},
//   	},
//   	Filter: &FilterProperty{
//   		RetainAtLeast: jsii.Number(123),
//   		Type: jsii.String("type"),
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html
//
type CfnLifecyclePolicyPropsMixin_PolicyDetailProperty struct {
	// Configuration details for the policy action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html#cfn-imagebuilder-lifecyclepolicy-policydetail-action
	//
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Additional rules to specify resources that should be exempt from policy actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html#cfn-imagebuilder-lifecyclepolicy-policydetail-exclusionrules
	//
	ExclusionRules interface{} `field:"optional" json:"exclusionRules" yaml:"exclusionRules"`
	// Specifies the resources that the lifecycle policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html#cfn-imagebuilder-lifecyclepolicy-policydetail-filter
	//
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}

