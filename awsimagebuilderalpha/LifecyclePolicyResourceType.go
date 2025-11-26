package awsimagebuilderalpha


// The resource type which the lifecycle policy is applied to.
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
type LifecyclePolicyResourceType string

const (
	// Indicates the policy applies to AMI-based images.
	// Experimental.
	LifecyclePolicyResourceType_AMI_IMAGE LifecyclePolicyResourceType = "AMI_IMAGE"
	// Indicates the policy applies to container images.
	// Experimental.
	LifecyclePolicyResourceType_CONTAINER_IMAGE LifecyclePolicyResourceType = "CONTAINER_IMAGE"
)

