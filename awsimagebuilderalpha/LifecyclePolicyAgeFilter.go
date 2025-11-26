package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// The age-based filtering to apply in a lifecycle policy rule.
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
type LifecyclePolicyAgeFilter struct {
	// The minimum age of the resource to filter.
	//
	// The provided duration will be rounded up to the nearest
	// day/week/month/year value.
	// Experimental.
	Age awscdk.Duration `field:"required" json:"age" yaml:"age"`
	// For age-based filters, the number of EC2 Image Builder images to keep on hand once the rule is applied.
	//
	// The value
	// must be between 1 and 10.
	// Default: 0.
	//
	// Experimental.
	RetainAtLeast *float64 `field:"optional" json:"retainAtLeast" yaml:"retainAtLeast"`
}

