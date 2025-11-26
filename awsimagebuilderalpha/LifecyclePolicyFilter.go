package awsimagebuilderalpha


// The resource filtering to apply in the lifecycle policy rule.
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
type LifecyclePolicyFilter struct {
	// The resource age filter to apply in the lifecycle policy rule.
	// Default: - none if a count filter is provided. Otherwise, an age filter is required.
	//
	// Experimental.
	AgeFilter *LifecyclePolicyAgeFilter `field:"optional" json:"ageFilter" yaml:"ageFilter"`
	// The resource count filter to apply in the lifecycle policy rule.
	// Default: - none if an age filter is provided. Otherwise, a count filter is required.
	//
	// Experimental.
	CountFilter *LifecyclePolicyCountFilter `field:"optional" json:"countFilter" yaml:"countFilter"`
}

