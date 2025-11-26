package awsimagebuilderalpha


// The count-based filtering to apply in a lifecycle policy rule.
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
type LifecyclePolicyCountFilter struct {
	// The minimum number of resources to keep on hand as part of resource filtering.
	// Experimental.
	Count *float64 `field:"required" json:"count" yaml:"count"`
}

