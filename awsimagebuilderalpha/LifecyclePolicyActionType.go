package awsimagebuilderalpha


// The action to perform on the resources which the policy applies to.
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
type LifecyclePolicyActionType string

const (
	// Indicates that the rule should delete the resource when it is applied.
	// Experimental.
	LifecyclePolicyActionType_DELETE LifecyclePolicyActionType = "DELETE"
	// Indicates that the rule should deprecate the resource when it is applied.
	// Experimental.
	LifecyclePolicyActionType_DEPRECATE LifecyclePolicyActionType = "DEPRECATE"
	// Indicates that the rule should disable the resource when it is applied.
	// Experimental.
	LifecyclePolicyActionType_DISABLE LifecyclePolicyActionType = "DISABLE"
)

