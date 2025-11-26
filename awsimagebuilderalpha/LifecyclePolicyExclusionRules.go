package awsimagebuilderalpha


// The rules to apply for excluding resources from the lifecycle policy rule.
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
type LifecyclePolicyExclusionRules struct {
	// The rules to apply for excluding AMIs from the lifecycle policy rule.
	// Default: - no exclusion rules are applied on the AMI.
	//
	// Experimental.
	AmiExclusionRules *LifecyclePolicyAmiExclusionRules `field:"optional" json:"amiExclusionRules" yaml:"amiExclusionRules"`
	// The rules to apply for excluding EC2 Image Builder images from the lifecycle policy rule.
	// Default: - no exclusion rules are applied on the image.
	//
	// Experimental.
	ImageExclusionRules *LifecyclePolicyImageExclusionRules `field:"optional" json:"imageExclusionRules" yaml:"imageExclusionRules"`
}

