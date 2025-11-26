package awsimagebuilderalpha


// The rules to apply for excluding EC2 Image Builder images from the lifecycle policy rule.
//
// Example:
//   excludeImagesPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("ExcludeImagesPolicy"), &LifecyclePolicyProps{
//   	ResourceType: imagebuilder.LifecyclePolicyResourceType_CONTAINER_IMAGE,
//   	Details: []LifecyclePolicyDetail{
//   		&LifecyclePolicyDetail{
//   			Action: &LifecyclePolicyAction{
//   				Type: imagebuilder.LifecyclePolicyActionType_DELETE,
//   			},
//   			Filter: &LifecyclePolicyFilter{
//   				CountFilter: &LifecyclePolicyCountFilter{
//   					Count: jsii.Number(20),
//   				},
//   			},
//   			ExclusionRules: &LifecyclePolicyExclusionRules{
//   				ImageExclusionRules: &LifecyclePolicyImageExclusionRules{
//   					Tags: map[string]*string{
//   						"DoNotDelete": jsii.String("true"),
//   						"Critical": jsii.String("baseline"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ResourceSelection: &LifecyclePolicyResourceSelection{
//   		Tags: map[string]*string{
//   			"Application": jsii.String("frontend"),
//   		},
//   	},
//   })
//
// Experimental.
type LifecyclePolicyImageExclusionRules struct {
	// Excludes EC2 Image Builder images with any of the provided tags from the lifecycle policy rule.
	// Experimental.
	Tags *map[string]*string `field:"required" json:"tags" yaml:"tags"`
}

