package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The rules to apply for excluding AMIs from the lifecycle policy rule.
//
// Example:
//   excludeAmisPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("ExcludeAmisPolicy"), &LifecyclePolicyProps{
//   	ResourceType: imagebuilder.LifecyclePolicyResourceType_AMI_IMAGE,
//   	Details: []LifecyclePolicyDetail{
//   		&LifecyclePolicyDetail{
//   			Action: &LifecyclePolicyAction{
//   				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
//   			},
//   			Filter: &LifecyclePolicyFilter{
//   				AgeFilter: &LifecyclePolicyAgeFilter{
//   					Age: awscdk.Duration_Days(jsii.Number(30)),
//   				},
//   			},
//   			ExclusionRules: &LifecyclePolicyExclusionRules{
//   				AmiExclusionRules: &LifecyclePolicyAmiExclusionRules{
//   					IsPublic: jsii.Boolean(true),
//   					 // Exclude public AMIs
//   					LastLaunched: awscdk.Duration_*Days(jsii.Number(7)),
//   					 // Exclude AMIs launched in last 7 days
//   					Regions: []*string{
//   						jsii.String("us-west-2"),
//   						jsii.String("eu-west-1"),
//   					},
//   					 // Exclude AMIs in specific regions
//   					SharedAccounts: []*string{
//   						jsii.String("123456789012"),
//   					},
//   					 // Exclude AMIs shared with specific accounts
//   					Tags: map[string]*string{
//   						"Protected": jsii.String("true"),
//   						"Environment": jsii.String("production"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ResourceSelection: &LifecyclePolicyResourceSelection{
//   		Tags: map[string]*string{
//   			"Team": jsii.String("infrastructure"),
//   		},
//   	},
//   })
//
// Experimental.
type LifecyclePolicyAmiExclusionRules struct {
	// Excludes public AMIs from the lifecycle policy rule if true.
	// Default: false.
	//
	// Experimental.
	IsPublic *bool `field:"optional" json:"isPublic" yaml:"isPublic"`
	// Excludes AMIs which were launched from within the provided duration.
	// Default: None.
	//
	// Experimental.
	LastLaunched awscdk.Duration `field:"optional" json:"lastLaunched" yaml:"lastLaunched"`
	// Excludes AMIs which reside in any of the provided regions.
	// Default: None.
	//
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Excludes AMIs which are shared with any of the provided shared accounts.
	// Default: None.
	//
	// Experimental.
	SharedAccounts *[]*string `field:"optional" json:"sharedAccounts" yaml:"sharedAccounts"`
	// Excludes AMIs which have any of the provided tags applied to it.
	// Default: None.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

