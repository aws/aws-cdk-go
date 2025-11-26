package awsimagebuilderalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a Lifecycle Policy resource.
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
type LifecyclePolicyProps struct {
	// Configuration details for the lifecycle policy rules.
	// Experimental.
	Details *[]*LifecyclePolicyDetail `field:"required" json:"details" yaml:"details"`
	// Selection criteria for the resources that the lifecycle policy applies to.
	// Experimental.
	ResourceSelection *LifecyclePolicyResourceSelection `field:"required" json:"resourceSelection" yaml:"resourceSelection"`
	// The type of Image Builder resource that the lifecycle policy applies to.
	// Experimental.
	ResourceType LifecyclePolicyResourceType `field:"required" json:"resourceType" yaml:"resourceType"`
	// The description of the lifecycle policy.
	// Default: None.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The execution role that grants Image Builder access to run lifecycle actions.
	//
	// By default, an execution role will be created with the minimal permissions needed to execute the lifecycle policy
	// actions.
	// Default: - an execution role will be generated.
	//
	// Experimental.
	ExecutionRole awsiam.IRole `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The name of the lifecycle policy.
	// Default: - a name is generated.
	//
	// Experimental.
	LifecyclePolicyName *string `field:"optional" json:"lifecyclePolicyName" yaml:"lifecyclePolicyName"`
	// The status of the lifecycle policy.
	// Default: LifecyclePolicyStatus.ENABLED
	//
	// Experimental.
	Status LifecyclePolicyStatus `field:"optional" json:"status" yaml:"status"`
	// The tags to apply to the lifecycle policy.
	// Default: - none.
	//
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

