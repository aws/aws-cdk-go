package awsimagebuilder


// The configuration details for a lifecycle policy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyDetailProperty := &PolicyDetailProperty{
//   	Action: &ActionProperty{
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		IncludeResources: &IncludeResourcesProperty{
//   			Amis: jsii.Boolean(false),
//   			Containers: jsii.Boolean(false),
//   			Snapshots: jsii.Boolean(false),
//   		},
//   	},
//   	Filter: &FilterProperty{
//   		Type: jsii.String("type"),
//   		Value: jsii.Number(123),
//
//   		// the properties below are optional
//   		RetainAtLeast: jsii.Number(123),
//   		Unit: jsii.String("unit"),
//   	},
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html
//
type CfnLifecyclePolicy_PolicyDetailProperty struct {
	// Configuration details for the policy action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html#cfn-imagebuilder-lifecyclepolicy-policydetail-action
	//
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Specifies the resources that the lifecycle policy applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html#cfn-imagebuilder-lifecyclepolicy-policydetail-filter
	//
	Filter interface{} `field:"required" json:"filter" yaml:"filter"`
	// Additional rules to specify resources that should be exempt from policy actions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html#cfn-imagebuilder-lifecyclepolicy-policydetail-exclusionrules
	//
	ExclusionRules interface{} `field:"optional" json:"exclusionRules" yaml:"exclusionRules"`
}

