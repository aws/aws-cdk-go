package awsimagebuilderalpha


// The action to perform in the lifecycle policy rule.
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
type LifecyclePolicyAction struct {
	// The action to perform on the resources selected in the lifecycle policy rule.
	// Experimental.
	Type LifecyclePolicyActionType `field:"required" json:"type" yaml:"type"`
	// Whether to include AMIs in the scope of the lifecycle rule.
	// Default: - true for AMI-based policies, false otherwise.
	//
	// Experimental.
	IncludeAmis *bool `field:"optional" json:"includeAmis" yaml:"includeAmis"`
	// Whether to include containers in the scope of the lifecycle rule.
	// Default: - true for container-based policies, false otherwise.
	//
	// Experimental.
	IncludeContainers *bool `field:"optional" json:"includeContainers" yaml:"includeContainers"`
	// Whether to include snapshots in the scope of the lifecycle rule.
	// Default: - true for AMI-based policies, false otherwise.
	//
	// Experimental.
	IncludeSnapshots *bool `field:"optional" json:"includeSnapshots" yaml:"includeSnapshots"`
}

