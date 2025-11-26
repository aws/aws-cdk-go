package awsimagebuilderalpha


// Configuration details for the lifecycle policy rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lifecyclePolicyDetail := &LifecyclePolicyDetail{
//   	Action: &LifecyclePolicyAction{
//   		Type: imagebuilder_alpha.LifecyclePolicyActionType_DELETE,
//
//   		// the properties below are optional
//   		IncludeAmis: jsii.Boolean(false),
//   		IncludeContainers: jsii.Boolean(false),
//   		IncludeSnapshots: jsii.Boolean(false),
//   	},
//   	Filter: &LifecyclePolicyFilter{
//   		AgeFilter: &LifecyclePolicyAgeFilter{
//   			Age: cdk.Duration_Minutes(jsii.Number(30)),
//
//   			// the properties below are optional
//   			RetainAtLeast: jsii.Number(123),
//   		},
//   		CountFilter: &LifecyclePolicyCountFilter{
//   			Count: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	ExclusionRules: &LifecyclePolicyExclusionRules{
//   		AmiExclusionRules: &LifecyclePolicyAmiExclusionRules{
//   			IsPublic: jsii.Boolean(false),
//   			LastLaunched: cdk.Duration_*Minutes(jsii.Number(30)),
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   			SharedAccounts: []*string{
//   				jsii.String("sharedAccounts"),
//   			},
//   			Tags: map[string]*string{
//   				"tagsKey": jsii.String("tags"),
//   			},
//   		},
//   		ImageExclusionRules: &LifecyclePolicyImageExclusionRules{
//   			Tags: map[string]*string{
//   				"tagsKey": jsii.String("tags"),
//   			},
//   		},
//   	},
//   }
//
// Experimental.
type LifecyclePolicyDetail struct {
	// The action to perform in the lifecycle policy rule.
	// Experimental.
	Action *LifecyclePolicyAction `field:"required" json:"action" yaml:"action"`
	// The resource filtering to apply in the lifecycle policy rule.
	// Experimental.
	Filter *LifecyclePolicyFilter `field:"required" json:"filter" yaml:"filter"`
	// The rules to apply for excluding resources from the lifecycle policy rule.
	// Default: - no exclusion rules are applied on any resource.
	//
	// Experimental.
	ExclusionRules *LifecyclePolicyExclusionRules `field:"optional" json:"exclusionRules" yaml:"exclusionRules"`
}

