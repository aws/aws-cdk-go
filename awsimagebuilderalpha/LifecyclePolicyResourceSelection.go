package awsimagebuilderalpha


// Selection criteria for the resources that the lifecycle policy applies to.
//
// Example:
//   disabledPolicy := imagebuilder.NewLifecyclePolicy(this, jsii.String("DisabledPolicy"), &LifecyclePolicyProps{
//   	LifecyclePolicyName: jsii.String("my-disabled-policy"),
//   	Description: jsii.String("A lifecycle policy that is temporarily disabled"),
//   	Status: imagebuilder.LifecyclePolicyStatus_DISABLED,
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
//   		},
//   	},
//   	ResourceSelection: &LifecyclePolicyResourceSelection{
//   		Tags: map[string]*string{
//   			"Environment": jsii.String("testing"),
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"Owner": jsii.String("DevOps"),
//   		"CostCenter": jsii.String("Engineering"),
//   	},
//   })
//
// Experimental.
type LifecyclePolicyResourceSelection struct {
	// The list of image recipes or container recipes to apply the lifecycle policy to.
	// Default: - none if tag selections are provided. Otherwise, at least one recipe or tag selection must be provided
	//
	// Experimental.
	Recipes *[]IRecipeBase `field:"optional" json:"recipes" yaml:"recipes"`
	// Selects EC2 Image Builder images containing any of the provided tags.
	// Default: - none if recipe selections are provided. Otherwise, at least one recipe or tag selection must be provided
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

